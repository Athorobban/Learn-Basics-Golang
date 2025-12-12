package main

import "fmt"

func main() {
	siswa := map[string]string{
		"nama":  "Sandi",
		"kelas": "2A",
	}
	fmt.Println("nama siswa : ", siswa["nama"])
	fmt.Println("kelas siswa : ", siswa["kelas"])

	siswa["alamat"] = "Kendal"
	fmt.Println("alamat :",siswa["alamat"])
}