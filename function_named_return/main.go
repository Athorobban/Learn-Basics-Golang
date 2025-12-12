package main

import "fmt"

func tambah(a, b int) (hasil int) {
	hasil = a + b
	return
}
func main() {
	fmt.Println(tambah(5,6))
}