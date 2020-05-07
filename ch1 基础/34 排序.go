package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println(strs)

	ints := []int{76, 2, 5}
	sort.Ints(ints)
	fmt.Println("Ints:", ints)
	s := sort.IntsAreSorted(ints)
	fmt.Println("sorted", s)
}
