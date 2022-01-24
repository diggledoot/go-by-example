package main

import "fmt"

func print(nums ...int) {
	i := 0
	for _, num := range nums {
		i += num
	}
	fmt.Println(nums, i)
}

func main() {
	print(1, 2, 3, 4, 5, 6, 7, 8, 9)
	print(1, 2)

	nums := []int{1, 2, 3, 4, 5}
	print(nums...)
}
