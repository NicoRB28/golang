package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//goroutine arranca y no se la espera, por eso el programa termina sin imprimir nada
	//no llega a hacer nada porque despues de la funcion no se hace nada mas
	go boring("hola: ")

}

func boring(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	}
}
