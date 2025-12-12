package main

import "fmt"

func main() {
	arr := [6]int{10, 20, 30, 40, 50, 60}

	s1 := arr[:] // ini akan membuat slice yang mencakup seluruh array
	fmt.Println("ini adalah s1 : ",s1)

	s2 := arr[:4] // ini akan membuat slice dari index 0 sampai index 3
	fmt.Println("ini adalah s2 : ",s2)

	s3 := arr[3:] // ini akan membuat slice dari index 3 sampai akhir array
	fmt.Println("ini adalah s3 : ",s3)

	s4 := arr[1:5] // ini akan membuat slice dari index 1 sampai index 4
	fmt.Println("ini adalah s4 : ",s4)
}