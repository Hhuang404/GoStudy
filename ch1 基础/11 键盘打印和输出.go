package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
输入和输出
fmt 包：输入和输出

输出：
	Print() // 打印
	Printf() // 格式化打印
	Println() // 打印之后换行
*/

func main() {
	// 输入
	var x int
	var y float64
	fmt.Println("请输入整数 和 浮点数")
	fmt.Scanln(&x, &y)
	fmt.Println(x, y)

	//使用 io 输入

	fmt.Println("请输入一个字符串：")
	reader := bufio.NewReader(os.Stdin)
	s1, _ := reader.ReadString('\n')
	fmt.Println("读到得数据：", s1)

}
