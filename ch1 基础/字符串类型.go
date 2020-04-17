package main

import (
	"fmt"
)

/**
字符串:
	1. 概念: 多个 byte 的集合, 理解为一个字符序列
	2。 语法: 使用双引号
		"abc", "hello","A"
*/
func main() {
	var s1 string = "huanghao"
	fmt.Printf("%T, %s \n", s1, s1)
	s2 := 'A'
	s3 := `C`
	fmt.Println(s2, s3) // 65 C
	s4 := '中'
	fmt.Printf("%T : %d,%c,%q", s4, s4, s4, s4)

}
