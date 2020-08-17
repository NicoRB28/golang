package main

import "fmt"

func main() {

	c := make(chan int)
	//deadlock
	c <- 42

	fmt.Println(<-c)

}
