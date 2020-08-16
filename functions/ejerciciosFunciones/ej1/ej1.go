package main

import "fmt"

func main() {

	fmt.Println("foo =>", foo())

	dia, mes := bar()
	fmt.Println("dia => ", dia, "mes => ", mes)

}

func foo() int {
	return 100
}

func bar() (int, string) {
	return 28, "marzo"
}
