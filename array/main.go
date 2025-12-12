package main

import "fmt"

func main() {
	var angka [3]int = [3]int{10, 20, 30}
	fmt.Println(angka)
	fmt.Println(angka[0])
	angka[1] = 50
	fmt.Println("ini adalah value index ke-1 : ", angka[1])
}