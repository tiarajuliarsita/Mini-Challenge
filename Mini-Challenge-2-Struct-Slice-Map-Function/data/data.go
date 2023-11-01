package data

import (
	"fmt"
	"os"
	"strconv"
)

type Member struct {
	ID        int //no absen
	Nama      string
	Pekerjaan string
	Alamat    string
	Alasan    string
}

func dataMember() []Member {
	members := []Member{
		{ID: 1, Nama: "tiara", Pekerjaan: "DevOps", Alamat: "palu", Alasan: "menyukai hal baru"},
		{ID: 2, Nama: "jennie", Pekerjaan: "Backend", Alamat: "soul", Alasan: "tertarik pada teknologi"},
		{ID: 3, Nama: "lisa", Pekerjaan: "Frontend", Alamat: "Tokyo", Alasan: "menyukai animasi"},
		{ID: 4, Nama: "rose", Pekerjaan: "Data Analyst", Alamat: "Bangkok", Alasan: "senang menganalisis data"},
		{ID: 5, Nama: "jisoo", Pekerjaan: "UI/UX Designer", Alamat: "kuala lumpur", Alasan: "senang mendesain"},
	}
	return members
}


func FindMember() {

	members := dataMember()
	if len(os.Args) < 2 {
		fmt.Print("masukkan nama atau nomor absen member\ncontoh : go run main.go (nama member) atau go run main.go (no absen)")
		return
	}

	input := os.Args[1]
	AbsensInt, _ := strconv.Atoi(input)
	for _, v := range members {
		if v.Nama == input {
			fmt.Printf(" ID : %d\n Nama : %s\n Alamat: %s\n Pekerjaan: %s\n Alasan: %s\n \n",
				v.ID, v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)
			return
		}
	}
	for _, v := range members {
		if v.ID == AbsensInt {
			fmt.Printf(" ID : %d\n Nama : %s\n Alamat: %s\n Pekerjaan: %s\n Alasan: %s\n \n",
				v.ID, v.Nama, v.Alamat, v.Pekerjaan, v.Alasan)
			return
		}
	}
	fmt.Println("tidak ada member dengan nama atau nomor absen : ", input)

}
