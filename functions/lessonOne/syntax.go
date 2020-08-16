package main

import (
	"fmt"
)

func main() {

	foo()

	bar("nicolas")

	s1 := woo("Nico")
	fmt.Println(s1)

	s2, s3 := doble("hola")
	fmt.Println(s2, " ", s3)
}

//func (r receiver) identifier(parameters)(return(s)){....}

func foo() {
	fmt.Println("hello from foo")
}

//los parametros son por VALOR
func bar(s string) {
	fmt.Println("hello ", s)
}

//luego del (con los parametros) se pone el tipo de dato que se retorna
func woo(s string) string {
	return fmt.Sprint("hello from woo, ", s)
}

//puedo retornar mas de un valor
func doble(s string) (string, int) {
	return s, 10
}
