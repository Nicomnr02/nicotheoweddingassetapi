package handler

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func Handler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	carouselImagePage0 := []string{
		"invitation-asset/1. Loading Spinner Photos/1.webp",
		"invitation-asset/1. Loading Spinner Photos/2.webp",
		"invitation-asset/1. Loading Spinner Photos/3.webp",
		"invitation-asset/1. Loading Spinner Photos/4.webp",
		"invitation-asset/1. Loading Spinner Photos/5.webp",
	}

	carouselImagePage1 := []string{
		"invitation-asset/2. Introduction Photos/1.webp",
		"invitation-asset/2. Introduction Photos/2.webp",
		"invitation-asset/2. Introduction Photos/3.webp",
		"invitation-asset/2. Introduction Photos/4.webp",
		"invitation-asset/2. Introduction Photos/5.webp",
	}

	bgImagePage3 := []string{
		"invitation-asset/3. Bible Verses Photo/1.webp",
	}

	bgImagePage4 := []string{
		"invitation-asset/4. Groom Photo/1.webp",
	}

	bgImagePage5 := []string{
		"invitation-asset/5. Bride Photo/1.webp",
	}

	bgImagePage6 := []string{
		"invitation-asset/6. Love Journey Photo/1.webp",
	}

	bgImagePage7 := []string{
		"invitation-asset/7. Save The Date Photo/1.webp",
	}

	bgImagePage8 := []string{
		"invitation-asset/8. Marriage Countdown Photo/1.webp",
	}

	bgImagePage9 := []string{
		"invitation-asset/9. Wishes Form Photo/1.webp",
	}

	bgImagePage10 := []string{
		"invitation-asset/10. Wishes List Photo/1.webp",
	}

	bgImagePage11 := []string{
		"invitation-asset/11. Wedding Gift Photo/1.webp",
	}

	bgImagePage12 := []string{
		"invitation-asset/12. Prewedding Collection Photos/1.webp",
		"invitation-asset/12. Prewedding Collection Photos/2.webp",
		"invitation-asset/12. Prewedding Collection Photos/3.webp",
		"invitation-asset/12. Prewedding Collection Photos/4.webp",
		"invitation-asset/12. Prewedding Collection Photos/5.webp",
		"invitation-asset/12. Prewedding Collection Photos/6.webp",
		"invitation-asset/12. Prewedding Collection Photos/7.webp",
		"invitation-asset/12. Prewedding Collection Photos/8.webp",
	}

	bgImagePage14 := []string{
		"invitation-asset/14. Thanking Photo/1.jpeg",
	}

	bgMusic := "invitation-asset/bg_music/household_of_faith.m4a"

	introVideo := "invitation-asset/intro_video/intro_video.mp4"

	page1Video := "invitation-asset/page_1_video/page_1_video.mp4"

	accessKey := os.Getenv("ACCESS_KEY")
	secretKey := os.Getenv("SECRET_KEY")
	bucket := "nico-wedding"

	// Initialize AWS S3 client but point it to GCS
	sess := session.Must(session.NewSession(&aws.Config{
		Region:           aws.String("auto"), // region is ignored but required
		Endpoint:         aws.String("https://storage.googleapis.com"),
		S3ForcePathStyle: aws.Bool(true),
		Credentials:      credentials.NewStaticCredentials(accessKey, secretKey, ""),
	}))

	svc := s3.New(sess)

	var sign = func(object string) string {
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(object),
		})

		// Generate pre-signed URL
		url, err := req.Presign(72 * time.Hour)
		if err != nil {
			http.Error(w, "Failed to sign data", http.StatusInternalServerError)
		}
		return url
	}

	response := struct {
		CarouselImagePage0 []string `json:"carousel_image_page_0"`
		CarouselImagePage1 []string `json:"carousel_image_page_1"`
		BGImagePage3       []string `json:"bg_image_page_3"`
		BGImagePage4       []string `json:"bg_image_page_4"`
		BGImagePage5       []string `json:"bg_image_page_5"`
		BGImagePage6       []string `json:"bg_image_page_6"`
		BGImagePage7       []string `json:"bg_image_page_7"`
		BGImagePage8       []string `json:"bg_image_page_8"`
		BGImagePage9       []string `json:"bg_image_page_9"`
		BGImagePage10      []string `json:"bg_image_page_10"`
		BGImagePage11      []string `json:"bg_image_page_11"`
		BGImagePage12      []string `json:"bg_image_page_12"` //prewed collection
		BGImagePage14      []string `json:"bg_image_page_14"`
		BGMusic            string   `json:"bg_music"`

		IntroVideo string `json:"intro_video"`
		Page1Video string `json:"page_1_video"`
	}{}

	for _, obj := range carouselImagePage0 {
		response.CarouselImagePage0 = append(response.CarouselImagePage0, sign(obj))
	}
	for _, obj := range carouselImagePage1 {
		response.CarouselImagePage1 = append(response.CarouselImagePage1, sign(obj))
	}

	for _, obj := range bgImagePage3 {
		response.BGImagePage3 = append(response.BGImagePage3, sign(obj))
	}
	for _, obj := range bgImagePage4 {
		response.BGImagePage4 = append(response.BGImagePage4, sign(obj))
	}
	for _, obj := range bgImagePage5 {
		response.BGImagePage5 = append(response.BGImagePage5, sign(obj))
	}
	for _, obj := range bgImagePage6 {
		response.BGImagePage6 = append(response.BGImagePage6, sign(obj))
	}
	for _, obj := range bgImagePage7 {
		response.BGImagePage7 = append(response.BGImagePage7, sign(obj))
	}
	for _, obj := range bgImagePage8 {
		response.BGImagePage8 = append(response.BGImagePage8, sign(obj))
	}
	for _, obj := range bgImagePage9 {
		response.BGImagePage9 = append(response.BGImagePage9, sign(obj))
	}
	for _, obj := range bgImagePage10 {
		response.BGImagePage10 = append(response.BGImagePage10, sign(obj))
	}
	for _, obj := range bgImagePage11 {
		response.BGImagePage11 = append(response.BGImagePage11, sign(obj))
	}
	for _, obj := range bgImagePage12 {
		response.BGImagePage12 = append(response.BGImagePage12, sign(obj))
	}
	for _, obj := range bgImagePage14 {
		response.BGImagePage14 = append(response.BGImagePage14, sign(obj))
	}

	response.BGMusic = sign(bgMusic)
	response.IntroVideo = sign(introVideo)
	response.Page1Video = sign(page1Video)

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
	}
}
