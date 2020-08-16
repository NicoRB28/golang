package main

import "fmt"

func main() {

	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	//como sum espera recibir uno o varios integers no puedo enviarle un slice de numeros, tengo que desarmarlo
	// y mandarle los numeros entonces puedo "desarmar" el slice con los ...
	x := sum(xi...)
	fmt.Println(x)

}

func sum(x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}

	return sum
}
