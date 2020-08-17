package main

import (
	"fmt"
	"runtime"
	"sync"
)

//golang.org/doc/effective_go.html#concurrency

//la funcion init se ejecuta siempre antes del main la idea es usarla para setear todo antes de arrancar
//func init() {
//go foo()
//}

var ws sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	//esto lo secuencializa mas que nada, primero se va a terminar bar y despues foo
	ws.Add(1) //-->hay una tarea a la que esperar

	go foo()
	bar()

	fmt.Println("CPUs\t", runtime.NumCPU())
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

	ws.Wait() //-->si llegas aca espera a que te den la señal de finalizado
	fmt.Println("Goroutines\t", runtime.NumGoroutine())

}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo: ", i)
	}

	ws.Done() //-->una vez que foo termina manda la señal
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar: ", i)
	}
}
