package main

import "fmt"

func main() {
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)

	// 通道缓冲
	mss := make(chan string, 2)
	mss <- "buffered"
	mss <- "channel"

	fmt.Println(<-mss)
	fmt.Println(<-mss)
}
