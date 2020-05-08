package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	rand.Seed(time.Now().Unix())
	fmt.Print(rand.Intn(100), ",")
	fmt.Print(rand.Intn(100))
	fmt.Println()
	fmt.Println(rand.Float64())
	fmt.Println((rand.Float64() * 5) + 5)
	fmt.Println((rand.Float64() * 5) + 5)

}
