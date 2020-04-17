package main

import (
	"fmt"
)

func main() {
	var name string
	name = "huanghao"
	fmt.Printf("名称是 %s\n", name)

	// var 变量名 数据类型 = 值
	var address string = "赣州市"
	fmt.Printf("地址是 %s\n", address)

	// 类型推断
	var number = "123544"
	fmt.Printf("类型自动推断 %s\n", number)

	// 简短声明
	sex := 1
	fmt.Printf("自动格式转换 当前格式为 %T ， sex =  %d  \n", sex, sex)

	//多变量声明
	var q, w, e int
	q = 1
	w = 2
	e = 3
	fmt.Println(q, w, e)

	// 同时声明多个 直接赋值
	var a, v, c = 100, 1.1, "123"

	fmt.Println(a, v, c)
	// 集合类型声明
	var (
		studentName = "123"
		studentSex  = "女"
	)
	fmt.Println(studentName, studentSex)

	var num int = 1
	fmt.Printf("num vale eq %d , the adress %p", num, &num)
}
