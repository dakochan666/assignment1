package main

import (
	"fmt"
	"os"
	"strconv"
)

type person struct {
	id        int
	nama      string
	alamat    string
	pekerjaan string
	alasan    string
}

var people = []person{
	{
		id:        1,
		nama:      "Rasyid",
		alamat:    "Tangerang",
		pekerjaan: "BackEnd Engineer",
		alasan:    "Suka",
	},
	{
		id:        2,
		nama:      "Adrian",
		alamat:    "Bekasi",
		pekerjaan: "BackEnd Engineer",
		alasan:    "Suka",
	},
	{
		id:        3,
		nama:      "Lius",
		alamat:    "Tangerang",
		pekerjaan: "BackEnd Engineer",
		alasan:    "Suka",
	},
	{
		id:        4,
		nama:      "Saipul",
		alamat:    "Ciamis",
		pekerjaan: "BackEnd Engineer",
		alasan:    "Suka",
	},
}

func main() {
	args := "tidak ada"
	if len(os.Args) > 1 {
		args = os.Args[1]
	} else {
		fmt.Println("mohon masukkan id yang ingin dicari")
		return
	}

	intArgs, err := strconv.Atoi(args)
	if err != nil {
		fmt.Println("terjadi kesalahan")
		return
	}

	if intArgs > len(people) || intArgs < 1 {
		fmt.Println("Data tidak ditemukan")
		return
	}

	for _, data := range people {
		if data.id == intArgs {
			fmt.Println("Nama :", data.nama)
			fmt.Println("Alamat :", data.alamat)
			fmt.Println("Pekerjaan :", data.pekerjaan)
			fmt.Println("Alasan memilih kelas Golang :", data.alasan)
		}
	}
}
