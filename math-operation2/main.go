package main

import "fmt"

func main() {
	x := 15

	x = 5
	fmt.Println(x)

	x += 10
	fmt.Println("x += 10 :", x)

	x -= 3
	fmt.Println("x -= 3 :", x)

	x *= 2
	fmt.Println("x *= 2 :", x)

	x /= 4
	fmt.Println("x /= 4 :", x)

	x %= 6
	fmt.Println("x %= 6 :", x)
}