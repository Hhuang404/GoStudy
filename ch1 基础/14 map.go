package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 7
	fmt.Println(m)

	v := m["k1"]
	fmt.Println(v)
	fmt.Println("len:", len(m))

	delete(m, "k1")
	fmt.Println(m)

	//当从一个 map 中取值时，可选的第二返回值指示这个键是否在这个 map 中。这可以用来消除键不存在和键有零值，像 0 或者 "" 而产生的歧义
	_, prs := m["k3"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)

}
