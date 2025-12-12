package main

import "fmt"

func main() {
	//deklarasi string menggunakan kutip ganda
	nama := "Robban"
	alamat := "Jl. Merdeka No. 10"

	//deklarasi string menggunakan backtick
	kalimat := `Halo, nama saya Robban.`

	//menampilkan nilai string
	fmt.Println("Nama :", nama)
	fmt.Println("Alamat :", alamat)
	fmt.Println("Kalimat :\n", kalimat)
}