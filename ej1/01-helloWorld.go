package main

import "fmt"

func main() {

	var name string
	fmt.Println("what's your name?:")
	fmt.Scan(&name)
	fmt.Println("hello ", name)
}
