package main

import "fmt"

func main() {

	c := make(chan int)

	go foo(c)

	for v := range c {
		fmt.Println(v)
	}

	//si no se pone el close en foo, el canal queda abierto entonces range quiere leer del canal pero foo ya no
	//carga mas datos en el canal ahi esta el deadlock
	//range saca valores del canal hasta que el canal se cierre

}

func foo(c chan<- int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
