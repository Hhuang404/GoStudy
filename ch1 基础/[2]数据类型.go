package main

import "fmt"

/**
Go语言的数据类型：
	1.基本数据类型
		布尔类型：bool
			取值：true，false
		数值类型：
			整数：int
				有符号: 最高位标识符号位，0 正数 1负数 ，其余位表示数值
					int8: (-128 到 127)
					int16: (-32768 到 32768)
					int32: (-2147483648 到 2147483648)
					int64: (-9223372036854775808 到 9223372036854775808)
				无符号：所有位表示数值
					uint8: (0 到 255)
					uint16: (0 到 65535)
					uint32: (0 到 4294967295)
					uint64: (0 到 18446744073709551615 )
				通常使用 int, unit 看操作系统是 64 还是32。
				byte: unint8
				rune: int32
			浮点: 生活中的小数
					float32, float64
			复数: complex
			字符串: string
	2.复合数据类型
		array, slice, map, function, pointer, struct, interface,channel..
*/
func main() {
	// 布尔类型
	var b1 bool = false
	fmt.Printf("%T : %t \n", b1, b1)
	b2 := true
	fmt.Printf("%T : %t \n", b2, b2)

	// 数值类型
	//var  i int8  =200    constant 200 overflows int8
	var i1 int = 100000
	fmt.Printf("%T : %d", i1, i1)
	//var i2 int64
	//i2 = i1 cannot use i1 (type int) as type int64 in assignment
	// 从语法的角度上看 int64 和 int 不是同一种类型
	var i3 rune = 150000
	fmt.Println(i3)

	// 浮点数
	var f1 float32 = 3.14
	var f2 float64 = 5.21616816
	fmt.Printf("%T : %f \n", f1, f1)
	fmt.Printf("%T : %.2f \n", f2, f2)
	fmt.Println(f1, f2)
	f3 := 0.166
	fmt.Printf("%T : %f", f3, f3)

}
