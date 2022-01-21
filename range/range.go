package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	//the _ is when we don't need the index, implicitly, index is supplied
	//we have to supply it or it will treat the num:=range nums as a condition
	for _, num := range nums {
		sum += num
	}

	fmt.Println("sum: ", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	//it can iterate over maps
	kvs := map[string]int{"a": 1, "b": 2}
	for k, v := range kvs {
		fmt.Printf("%s -> %d\n", k, v)
	}

	//iterate over keys of the map
	for k := range kvs {
		fmt.Println("key: ", k)
	}

	//can iterate over string
	for i, c := range "hello world" {
		fmt.Println(i, c)
	}
}
