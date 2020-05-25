package main

import "fmt"

func main() {

	var smallNumber int
	var largeNumber int

	fmt.Println("enter a small number:")
	fmt.Scan(&smallNumber)
	fmt.Println("enter a large number")
	fmt.Scan(&largeNumber)
	fmt.Printf("(%d div %d)remainder => %d", largeNumber, smallNumber, largeNumber%smallNumber)

}
