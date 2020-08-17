package main

import "fmt"

func main() {

	c := make(chan int)
	receiver := make(<-chan int) //solo recibe
	sender := make(chan<- int)   //solo envia

	fmt.Printf("%T\n", c)
	fmt.Printf("%T, (only receive)\n", receiver)
	fmt.Printf("%T (only send)\n", sender)
}
