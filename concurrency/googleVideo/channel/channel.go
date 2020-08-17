package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	//para generar un canal:
	//1.Declaracion e inicializacion:
	//var c chan int
	//c = make(chan int)
	//o 2. todo junto: c:= make(chan int)
	//sending to a channel: c <- dato
	//receiving from a channel: value:= <- c
	//cuando se ejecuta <-c quien lo haya hecho queda en espera ya sea de que el dato enviado sea recibido
	//como tambien que el dato sea enviado
	c := make(chan string)
	go boring("hola: ", c)
	for i := 0; i < 5; i++ {
		fmt.Printf("you say: %q\n", <-c)
	}
	fmt.Println("end")

}

func boring(msg string, c chan string) {
	for i := 0; ; i++ {
		c <- fmt.Sprintf("%s %d", msg, i)
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
	}
}
