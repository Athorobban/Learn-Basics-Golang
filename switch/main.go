package main

import "fmt"

func main() {
	hari := "Sabtu"

	switch hari {
	case "Senin", "Selasa", "Rabu", "Kamis", "Jumat":
		fmt.Println("Hari Kerja")
	case "Sabtu", "Minggu":
		fmt.Println("Hari Libur")
	default:
		fmt.Println("Hari tidak valid")	
	}

	// switch short statement
	switch panjang := len("Informatika"); {
	case panjang > 10:
		fmt.Println("Kata terlalu panjang")
	case panjang >= 5:
		fmt.Println("Kata sedang")
	default:
		fmt.Println("Kata terlalu pendek")
	}
}