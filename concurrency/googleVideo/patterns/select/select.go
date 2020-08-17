package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := fanIn(boring("Pepe"), boring("Juan"))
	for i := 0; i < 22; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("end")
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
				//si ningun canal esta disponible se bloquea hasta que alguno este, para evitar eso se puede agregar
				//el default:
			}
		}
	}()

	return c

}

func boring(msj string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msj, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
