package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("空数组 默认值 零", a)

	a[4] = 100
	fmt.Println("set", a)
	fmt.Println("get", a[4])

	b := [3]int{1, 2, 3}
	fmt.Println(b)
	var towD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			towD[i][j] = i + j
		}
	}
	fmt.Println("2d :", towD)
}
