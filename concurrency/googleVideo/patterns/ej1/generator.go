package main

import (
	"fmt"
	"math/rand"
	"time"
)

//una funcion que retorna un channel

func main() {

	c := boring("mensaje: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%s\n", <-c)
	}

	roberto := boring("Roberto")
	pepe := boring("Pepe")
	for j := 0; j < 4; j++ {
		fmt.Println(<-roberto)
		fmt.Println(<-pepe)
	}

	fmt.Println("end")

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
