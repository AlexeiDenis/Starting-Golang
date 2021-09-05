package main

import (
	"fmt"
	"time"
)

var (
	count int
)

func main() {

	iter := 1000
	for i := 0; i < iter; i++ {
		go increment()
	}
	time.Sleep(3 * time.Second)

	fmt.Println("salut: ", count)
}

func increment() {
	count++
}
