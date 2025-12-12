package main

import "fmt"

func jumlahAngka(angka ...int) int {
	total := 0
	for _, v := range angka {
		total += v
	}
	return total
}
func main() {
	fmt.Println(jumlahAngka(1, 2, 3))
	fmt.Println(jumlahAngka(10, 20, 30, 40, 50))
	fmt.Println(jumlahAngka())

}