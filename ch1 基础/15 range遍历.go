package main

import (
	"fmt"
)

func main() {

	nums := []int{2, 3, 4}
	sum := 0
	for num := range nums {
		sum += num
	}
	// index
	for index, value := range nums {
		fmt.Println(index, value)
	}

	kvs := map[string]uint8{"int": 123, "string": 234, "double": 3455}
	fmt.Println(kvs)
}
