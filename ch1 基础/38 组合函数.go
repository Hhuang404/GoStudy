package main

import (
	"fmt"
	"strings"
)

func index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

func include(vs []string, t string) bool {
	return index(vs, t) >= 0

}

func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

func all(vs []string, f func(string) bool) bool {
	for _, i2 := range vs {
		if !f(i2) {
			return false
		}
	}
	return true
}
func filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}
func mapl(vs []string, f func(string) string) []string {
	up := make([]string, len(vs))
	for i, v := range vs {
		up[i] = f(v)
	}
	return up
}

func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}
	fmt.Println(index(strs, "pear"))
	fmt.Println(include(strs, "grape"))
	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))
	fmt.Println(all(strs, func(s string) bool {
		return strings.HasPrefix(s, "p")
	}))
	fmt.Println(filter(strs, func(s string) bool {
		return strings.Contains(s, "e")
	}))
	fmt.Println(mapl(strs, strings.ToUpper))
}
