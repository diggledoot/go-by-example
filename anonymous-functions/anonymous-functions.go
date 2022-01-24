package main

import "fmt"

func intSeq() func() int {
	i := 0
	fmt.Printf("Outside call i is %d\n", i)
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())
}
