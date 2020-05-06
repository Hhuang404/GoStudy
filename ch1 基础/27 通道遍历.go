package main

import (
	"fmt"
)

func main() {
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	// 如果不关闭通道则会造成死锁 fatal error: all goroutines are asleep - deadlock!
	//close(queue)
	for elem := range queue {
		fmt.Println(elem)
	}
}
