package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(2, 2)
	fmt.Println("2+2", res)

	res = plusplus(2, 2, 2)
	fmt.Println("2+2+2 = ", res)
}
