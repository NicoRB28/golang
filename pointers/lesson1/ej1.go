package main

import "fmt"

func main() {

	a := 42
	fmt.Println("value:", a)
	fmt.Println("address:", &a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) //==> *int puntero a un int
	fmt.Println(*&a)       //==> el &a me esta dando la direccion apuntada y el * me da el valor almacenado en la dir apuntada

	var b *int = &a //==> aca *int es un tipo de dato puntero a int
	fmt.Println(b)
	fmt.Println(*b) //==>aca *b desreferenciar el puntero

	*b = 500
	fmt.Println("value:", a) //==> modifico el valor almacenado en a
}
