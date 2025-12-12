package main

import "fmt"

func main() {
	number := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Jumlah elemen : ", len(number))
	fmt.Println("Index Ke-1 :",number[1])
	number[2] = 20
	fmt.Println("Index Ke-2 setelah diubah :",number[2])

	fmt.Println("Ini adalah array :")

	for index, value := range number {
		fmt.Println("isi index ke : ", index," = ", value)
	}

	arr1 := [3]string{"apel", "jeruk", "mangga"}
	arr2 := [3]string{"bayam", "kangkung", "wortel"}

	fmt.Println("arr1 == arr2 : ", arr1 == arr2)
	fmt.Println("arr1 != arr2 : ", arr1 != arr2)
}