package main

import "fmt"

func half(a int) (int, bool) {

	b := (a / 2)
	c := a%2 == 0
	return b, c
}

func main() {

	var number int
	fmt.Println(`numero:`)
	fmt.Scan(&number)

	b, c := half(number)

	fmt.Printf("el numero: %d dividido 2 es igual a: %d y era par? %t", number, b, c)

}
