package main

import (
	"fmt"
)

/**
iota: 特殊的常量，可以被编译器自动修改的常量
		每当定义一个 const , iota 的初始值为 0
		每当定义一个常量，就会自动累加 1
		直到下一个 const 出现，清零
*/
func main() {
	const (
		a = iota  // iota 0
		b         // iota 1
		c = "123" // 123
		d         // 123 iota 2
	)
	const (
		q = iota
		w = iota
	)
	// 枚举
	const (
		MALE = iota
		FEMALE
		UNKNOWN
	)
	fmt.Println(a, b, c, d)
	fmt.Println("第二个常量出现 iota 清零 : ", q, w)
	fmt.Println(MALE, FEMALE, UNKNOWN)
}
