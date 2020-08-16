package main

import (
	"fmt"
)

func main() {
	sum := foo(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("el valor de la suma total desde el main es :", sum)
}

//esto me permite mandar varios parametros del tipo int, la notacion son los tres puntos ANTES del tipo ...int
func foo(x ...int) int {

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for _, value := range x {
		sum += value
	}

	fmt.Println("sum = ", sum)
	return sum

}
