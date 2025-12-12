package main

import "fmt"

func main() {
	var angaka1, angka2 int
	// input dari user
	fmt.Print("Masukkan angka pertama : ")
	fmt.Scanln(&angaka1)
	fmt.Print("Masukkan angka kedua : ")
	fmt.Scanln(&angka2)
	// perbandingan
	fmt.Println("\n=== Hasil Perbandingan ===")
	fmt.Printf("%d == %d ? %v\n", angaka1, angka2, angaka1 == angka2)
	fmt.Printf("%d != %d ? %v\n", angaka1, angka2, angaka1 != angka2)
	fmt.Printf("%d > %d ? %v\n", angaka1, angka2, angaka1 > angka2)
	fmt.Printf("%d < %d ? %v\n", angaka1, angka2, angaka1 < angka2)
	fmt.Printf("%d >= %d ? %v\n", angaka1, angka2, angaka1 >= angka2)
	fmt.Printf("%d <= %d ? %v\n", angaka1, angka2, angaka1 <= angka2)
}