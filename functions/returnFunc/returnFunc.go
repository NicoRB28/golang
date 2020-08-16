package main

import "fmt"

func main() {
	s1 := foo()
	fmt.Println(s1)

	s2 := bar()

	fmt.Printf("%T\n", s2)
	fmt.Println(s2())
	//esto tambien es valido:
	fmt.Println(bar()())
}

func foo() string {
	return "hello world"
}

//bar tiene como tipo de retorno un funcion que devuelve a su vez un int
//keyword-ID-parametros-tipo retornado
func bar() func() int {
	return func() int {
		return 451
	}
}
