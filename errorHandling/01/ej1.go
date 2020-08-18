package main

import "fmt"

func main() {

	//todo tipo que implemente el metodo Error()string es de tipo error
	//porque error es una iterface: type error interface{Error()string}

	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(n)
	fmt.Println(err)
}
