package main

import "fmt"

func main() {
	counter := 0
	c := make(chan int)

	go increment(c)
	for counter < 100 {
		counter += <-c
		fmt.Println("el contador va:", counter)
	}
	fmt.Printf("%d", counter)
}

func increment(c chan int) {
	for {
		c <- 1
	}
}
