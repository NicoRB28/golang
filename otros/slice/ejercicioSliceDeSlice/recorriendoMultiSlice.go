package main

import (
	"fmt"
)

func main() {

	sliceOne := []string{"James", "Bond", "Shaken, not stirred"}
	sliceTwo := []string{"Miss", "Moneypenny", "Hellooooo james"}
	multiSlice := [][]string{sliceOne, sliceTwo}

	//el caracter _ se usa para perder el valor que se devolvio en este caso el index porque no lo voy a usar
	for _, v := range multiSlice {
		for _, value := range v {
			fmt.Println(value)
		}
	}

}
