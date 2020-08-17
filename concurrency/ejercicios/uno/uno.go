package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {

	wg.Add(2)

	go foo()
	go bar()

	wg.Wait()
	fmt.Println("end")
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("hola")
	}

	wg.Done()

}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("mundo")
	}
	wg.Done()
}
