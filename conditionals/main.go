package main

import "fmt"

func main() {
	var nilai int
	fmt.Print("Masukkan nilai ujian : ")
	fmt.Scanln(&nilai)

	if nilai >= 90 {
		fmt.Println("Nilai Anda A (sangat baik).")
	} else if nilai >= 75 {
		fmt.Println("Nilai Anda B (baik).")
	} else if nilai >= 60 {
		fmt.Println("Nilai Anda C (cukup).")
	} else {
		fmt.Println("Nilai Anda D (perlu perbaikan).")
	}
	
}