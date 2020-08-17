package main

import (
	"fmt"
	"runtime"
)

func main() {

	c := make(chan int, 2) //esto es un buffer chan, permite almacenar un dato
	//no es recomendable el uso de buffers

	fmt.Println("GOROUTINS:", runtime.NumGoroutine())

	c <- 42
	c <- 54

	fmt.Println("GOROUTINS:", runtime.NumGoroutine())

	fmt.Println(<-c)
	fmt.Println(<-c)
}
