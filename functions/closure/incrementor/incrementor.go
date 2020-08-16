package main

import "fmt"

func main() {

	a := incrementor()
	b := incrementor()
	fmt.Println("con A:")
	fmt.Println("value: ", a())
	fmt.Println("value: ", a())
	fmt.Println("value: ", a())
	fmt.Println("value: ", a())
	fmt.Println("ahora con b:")
	fmt.Println("value: ", b())
	fmt.Println("value: ", b())
	fmt.Println("value: ", b())
	fmt.Println("value: ", b())
}

func incrementor() func() int {
	var x int
	//x es accesible dentro de la funcion anonima
	return func() int {
		x++
		return x
	}
}
