package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

// 多返回值
func vals() (int, int) {
	return 3, 4
}

// 变参函数
func sum(nums ...int) {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

// 闭包
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

// 递归
func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func main() {
	fmt.Println(plus(1, 2))
	fmt.Println(vals())
	sum(1, 2, 3, 4, 455, 6, 7)

	nexInt := intSeq()
	fmt.Println(nexInt())
	fmt.Println(nexInt())

	newInts := intSeq()
	fmt.Println(newInts())

	fmt.Println(fact(3))

}
