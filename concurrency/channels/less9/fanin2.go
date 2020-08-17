package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	c := fanIn(boring("joe"), boring("ann"))
	fmt.Println("goroutines", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			//le saco el dato a input y lo guardo en c
			c <- <-input1
		}
	}()

	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
