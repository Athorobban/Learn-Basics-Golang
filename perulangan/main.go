package main

import "fmt"

func main() {
	fmt.Println("=== Contoh Kombinasi FOR Loop di Bahasa Go ===")

	// 1. For Standar: mencetak angka 1 sampai 5
	fmt.Println("\n1. For standar (i = 1; i <= 5; i++):")
	for i := 1;i <= 5; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// 2. while-style: mencetak angka genap sampai 20
	fmt.Println("\n2. while-style (hanya kondisi):")
	j := 2
	for j <= 20 {
		fmt.Printf("%d ", j)
		j += 2
	}
	fmt.Println()

	// 3. infinite loop dengan break: menghitung mundur dari 10
	fmt.Println("\n3. Infinite loop dengan break:")
	k := 10
	for {
		fmt.Printf("%d ", k)
		k--
		if k == 0 {
			break // keluar dari loop ketika k mencapai 0
		}
	}
	fmt.Println()

	// 4. Range loop: iterasi elemen slice
	fmt.Println("\n4. Range loop untuk slice:")
	buah := []string{"apel", "pisang", "mangga"}
	for index, value := range buah {
		fmt.Printf("index %d: %s\n", index, value)
	}
}