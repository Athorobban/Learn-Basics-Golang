package main

import (
	"fmt"
	"strconv"
)

func main() {
	var score int = 100
	var text string = strconv.Itoa(score)//konversi int ke string

	fmt.Println("Nilai score :", text)
}