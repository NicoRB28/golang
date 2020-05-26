package main

import "fmt"

func half(a float64) (float64, bool) {

	b := (a / 2)
	c := int(a)%2 == 0
	return b, c
}

func main() {

	var number float64
	fmt.Println(`numero:`)
	fmt.Scan(&number)

	b, c := half(number)

	fmt.Printf("el numero: %d dividido 2 es igual a: %.2f y era par? %t", int(number), b, c)

}
