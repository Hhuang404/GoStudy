package main

import (
	"fmt"
)

/**
算术运算符： +，-，*，/，%，++，--

*/
func main() {
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d = %d \n", a, b, sum)

	sub := a - b
	fmt.Printf("%d - %d = %d \n", a, b, sub)

	mul := a * b
	fmt.Printf("%d * %d = %d \n", a, b, mul)

	div := a / b
	mod := a % b
	fmt.Printf("%d / %d = %d \n", a, b, div)
	fmt.Printf("%d %% %d = %d \n", a, b, mod)

}
