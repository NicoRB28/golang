package main

import (
	"fmt"
)

func main() {

	//como la funcion es anonima no tiene ID por eso para que funcione es necesario agregar () al final
	func() {
		fmt.Println("funcion anonima sin parametros")
	}()

	func(s string) {
		fmt.Println("soy la ", s, " funcion anonima esta vez con parametro")
	}("segunda")

	//tambien se puede almacenar la funcion en una variable
	a := func() {
		fmt.Println("esta funcion es anonima, no tiene parametros y esta almacenada en la var a")
	}

	b := func(s string) {
		fmt.Println("esta funcion es anonima, si tiene parametros: ", s, " y esta almacenada en b")
	}

	a()
	b("parametro de B")

}
