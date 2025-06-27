package main

import (
	"fmt"
	"nicotheowedding/api/whatsapp"
	"strings"
)

const (
	NomerMami   = "6285261919124"
	NomerSayang = "6281370905656"
)

func main() {

	wa := whatsapp.NewWhatsappClient()

	//! YANG UDAH DI COMMENT BRTI UDAH DI GENERATE
	list := []string{
		// "Pdt Boas Malango",
		// "Pdt. Beres Panjaitan",
		// "GroupAngkatan PTASN 99",
		// "Group Sehati Sejiwa",
		// "Group.Pasombu Sihol",
		// "Group Imanuel Ministry", // typo mungkin: Group Imanuel Ministry
		// "Group Ministry Doa Ibu",
		// "Group Pendeta Se WKB",
		// "kel Jufri si mangunsong/Ranni Si debang",
		// "Kel Erna Nababan",
		// "Kel.Frenly Purba",
		// "Kel.Gredy Purba",
		// "Kel Chandro Tobing",
		// "Kel Sitanggan/Rut br Manurung",

		// "GMAHK Kisaran",
		// "GMAHK labuhan ruku",
		// "GMAHK Meranti",
		// "Kel bpk Silalahi/br hotang",
		// "Kel.Bpk Situngkir",
		// "Kel besar Pinoppar Op Tambok Sihombing",
		// "J.Lingga (Poli)",
		// "Ibu B. Hasibuan",
		// "H.Marbun",
		// "I.Simanjuntak",

		// "Daniel Capah",

		// "Kel bpk L Silalahi/br Sihotang",
		// "Kel Sitinjak/br silalahi",
	}

	for _, item := range list {

		// Step 1: Ganti titik dengan spasi
		item = strings.ReplaceAll(item, ".", " ")

		// Step 2: Trim untuk menghilangkan spasi ekstra
		item = strings.TrimSpace(item)

		// Step 3: Ganti semua spasi menjadi %20
		item = strings.ReplaceAll(item, " ", "%20")

		err := wa.SendPersonalMessage(NomerMami, fmt.Sprintf(
			`
	Shalom dan salam hangat,
	
	Dengan penuh sukacita dan ucapan syukur kepada Tuhan Yesus Kristus, kami mengundang Bapak/Ibu/Saudara/i terkasih untuk hadir dan memberikan doa restu di hari bahagia kami, saat kami mengikat janji suci pernikahan:
		
	Nicolas & Theofani
	
	📆 Selasa, 8 Juli 2025
		
	✝ Ibadah Pemberkatan Pernikahan
	🕘 Pukul 09.00 WIB – Selesai
	📍 Gereja Masehi Advent Hari Ketujuh (GMAHK) Tanjungbalai
		
	💍 Acara Adat dan Resepsi
	🕛 Pukul 12.00 WIB – Selesai
	📍 Aula Pertemuan Katolik Teluk Ketapang, Tanjungbalai
		
	Kami percaya bahwa perjalanan ini adalah karya kasih Tuhan yang begitu indah. Dan pada hari itu, kami berharap dapat berbagi sukacita bersama Bapak/Ibu/Saudara/i yang telah menjadi bagian penting dalam hidup kami.
		
	Informasi lebih lengkap mengenai acara pernikahan kami dapat diakses melalui tautan berikut:
	🌐 https://nicotheoweddinginvitation.netlify.app/?guest_name=%s
		
	Kehadiran dan doa restu Anda merupakan anugerah besar bagi kami. Kiranya Tuhan senantiasa melimpahkan kesehatan, kebahagiaan, dan damai sejahtera bagi Bapak/Ibu/Saudara/i di setiap langkah kehidupan.
		
	Dengan kasih,
	Nico & Fani
		`, item))
		if err != nil {
			fmt.Println("fail to SendPersonalMessage because of " + err.Error())
		} else {
			fmt.Println("success to SendPersonalMessage of " + item)
		}
	}

}
