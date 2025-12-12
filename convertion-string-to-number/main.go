package main

import (
	"fmt"
	"strconv"
)

func main() {
	var text string = "100"
	number, err := strconv.Atoi(text) //konversi string ke int
	if err != nil {
		fmt.Println("Error konversi:", err)
	}else{
		fmt.Println("Nilai number :", number)
	}
}