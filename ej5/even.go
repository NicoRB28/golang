package main

import "fmt"

func main() {

	fmt.Println("se van a imprimir los pares entre 0 y 100:")
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
