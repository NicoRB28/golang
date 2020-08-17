package main

import "fmt"

func main() {

	c := make(chan<- int, 2) //al agregar la flecha en el make, genero que el canal sea solo para enviar
	//no puedo usar este canal para sacar datos
	d := make(<-chan int, 2) // la flecha antes del chan genera que el canal sea solo para recibir datos es decir
	//que puede sacar datos de el

	c <- 42
	c <- 54

	fmt.Println(<-d)
	fmt.Println(<-d)
	fmt.Println("-----")
	fmt.Printf("el tipo de la variable c es: %T\n", c)

}
