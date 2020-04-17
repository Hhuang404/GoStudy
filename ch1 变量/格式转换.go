package main

import (
	"fmt"
)

/**
类型转换
格式: Type(Value)
*/
func main() {
	var a int8 = 10

	var b int16

	b = int16(a)
	fmt.Println(b)

	f1 := 4.86
	var c int
	c = int(f1)
	fmt.Println(c)

	//bo := false
	//a = int8(bo) // cannot convert bo (type bool) to type int8
	//fmt.Println(a,bo)

	sum := f1 + 100
	fmt.Printf("%T: %f", sum, sum)
}
