package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println("emp：", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println(c)

	l := c[2:5]
	fmt.Println("获取2-5的数据", l)

	q := c[3:]
	fmt.Println("获取下标 3 及以后的数据", q)

	w := c[:4]
	fmt.Println("获取下标 4 以及以前的数据", w)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
