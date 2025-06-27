package whatsapp

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"log"
	"mime/multipart"
	"net/http"

	"golang.org/x/time/rate"
)

type Whatsapp interface {
	SendPersonalMessage(phoneNumber string, message string) error
	SendGroupMessage(groupID string, message string) error
}

type WhatsappClient struct {
	apiURL      string
	accessToken string
	httpClient  *http.Client
	limiter     *rate.Limiter
}

const (
	sendTextEndpoint = "/send/text"

	authHeader        = "Authorization"
	contentTypeHeader = "Content-Type"

	bearerAuthType = "Bearer "
)

func NewWhatsappClient() Whatsapp {
	rateLimitPerSecond := rate.Limit(10.0 / 5.0)
	burstSize := 10

	return &WhatsappClient{
		apiURL:      "https://go-wa.glvm.site/access",
		accessToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXQiOnsiamlkIjoiNjI4MjMyMjM5MzcyMCJ9LCJleHAiOjIxMTEwNDIzNDIsImlhdCI6MTc1MTA0MjM0Mn0.BtriYll0rmfOLPxG8Tf3YDw42x578iNW4DtXQLOaNhA",
		httpClient:  &http.Client{},
		limiter:     rate.NewLimiter(rateLimitPerSecond, burstSize),
	}
}

func (w *WhatsappClient) doSendRequest(chatID string, message string) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", w.apiURL, sendTextEndpoint)

	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)

	if err := writer.WriteField("msisdn", chatID); err != nil {
		return nil, fmt.Errorf("failed to write msisdn field: %w", err)
	}
	if err := writer.WriteField("message", message); err != nil {
		return nil, fmt.Errorf("failed to write message field: %w", err)
	}
	if err := writer.Close(); err != nil {
		return nil, fmt.Errorf("failed to close multipart writer: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, payload)
	if err != nil {
		log.Printf("Error creating send message request: %v", err)
		return nil, fmt.Errorf("failed to create send message request: %w", err)
	}

	token := w.accessToken
	if token == "" {
		log.Println("Attempted to send request with empty access token in doSendRequest")
		return nil, fmt.Errorf("internal error: access token is empty before sending request")
	}
	req.Header.Add(authHeader, bearerAuthType+token)
	req.Header.Set(contentTypeHeader, writer.FormDataContentType())

	res, err := w.httpClient.Do(req)
	if err != nil {
		log.Printf("Error performing send message request to %s: %v", chatID, err)
		return nil, fmt.Errorf("failed to perform send message request: %w", err)
	}

	return res, nil
}

func (w *WhatsappClient) sendMessage(chatID string, message string) error {
	if chatID == "" {
		return fmt.Errorf("chat ID cannot be empty")
	}
	if message == "" {
		return fmt.Errorf("message cannot be empty")
	}

	ctx := context.Background()
	if err := w.limiter.Wait(ctx); err != nil {
		log.Printf("Rate limiter wait failed for %s: %v", chatID, err)
		return fmt.Errorf("rate limiter wait failed: %w", err)
	}

	res, err := w.doSendRequest(chatID, message)
	if err != nil {
		log.Printf("Failed during send request for %s: %v", chatID, err)
		return fmt.Errorf("failed during send request: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		return nil
	} else if res.StatusCode == http.StatusUnauthorized { // 401 Unauthorized
		log.Printf("Send message failed with 401 Unauthorized for %s.", chatID)
		return fmt.Errorf("send message failed with 401 Unauthorized")
	} else {
		bodyBytes, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			log.Printf("Failed to send message to %s. Status: %d. Error reading body: %v", chatID, res.StatusCode, readErr)
			return fmt.Errorf("failed to send message to %s: API returned status %d (error reading body: %v)", chatID, res.StatusCode, readErr)
		} else {
			log.Printf("Failed to send message to %s. Status: %d. Body: %s", chatID, res.StatusCode, string(bodyBytes))
			return fmt.Errorf("failed to send message to %s: API returned status %d, body: %s", chatID, res.StatusCode, string(bodyBytes))
		}
	}
}

func (w *WhatsappClient) SendPersonalMessage(phoneNumber string, message string) error {
	if phoneNumber == "" {
		return fmt.Errorf("personal phone number cannot be empty")
	}
	return w.sendMessage(phoneNumber+"@c.us", message)
}

func (w *WhatsappClient) SendGroupMessage(groupID string, message string) error {
	if groupID == "" {
		return fmt.Errorf("group ID cannot be empty")
	}
	return w.sendMessage(groupID+"@g.us", message)
}
