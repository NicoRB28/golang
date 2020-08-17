package main

import "fmt"

func doSomething(num int) int {
	return num * 5
}

func main() {
	ch := make(chan int)

	go func() {
		//meto en el canal el resultado de la funcion doSomething
		ch <- doSomething(5)
	}()

	fmt.Println(<-ch) //aca saco del canal el resultado para imprimirlo

}
