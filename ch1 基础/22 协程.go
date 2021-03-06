package main

import "fmt"

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	f("direct")

	go f("huangh")

	go func(str string) {
		fmt.Println(str, ":", str)
	}("going")

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
