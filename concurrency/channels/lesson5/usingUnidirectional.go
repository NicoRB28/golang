package main

import "fmt"

func main() {

	//tengo un canal bidireccional
	c := make(chan int)

	//se lo mando a foo que tiene como parametro un canal unidireccional de solo envio
	go foo(c)
	//se lo mando a bar que tiene como parametro un canal unidireccional de solo recibir
	bar(c)
	//la transformacion de canales solo es posible de bidireccional o generico a unidireccional o especifico
	//no se puede transformar un canal unidireccional en uno bidireccional

	fmt.Println("end")
}

//send ( en esta funcion c es un canal solo para enviar)
func foo(c chan<- int) {
	c <- 28
}

//receive (en esta funcion c es un canal solo para recibir)
func bar(c <-chan int) {
	fmt.Println(<-c)
}
