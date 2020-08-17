package main

import (
	"fmt"
	"runtime"
)

func main() {

	c := make(chan int)

	fmt.Println("GOROUTINS:", runtime.NumGoroutine()) //en este punto hay una sola (main)

	go func() { c <- 42 }()

	fmt.Println("GOROUTINS:", runtime.NumGoroutine()) //en este punto hay dos main y la func literal
	fmt.Println(<-c)

}
