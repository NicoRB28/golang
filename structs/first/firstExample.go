package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {

	p1 := person{
		firstName: "Nicolas",
		lastName:  "Boschi",
		age:       28,
	}

	p2 := person{
		firstName: "Roberto",
		lastName:  "Bolaños",
		age:       90,
	}

	fmt.Println("primer nombre:", p1.firstName, " apellido: ", p1.lastName)
	fmt.Println("primer nombre:", p2.firstName, " apellido: ", p2.lastName)

	fmt.Println("tambien se puede hacer: ")
	fmt.Println(p1)
	fmt.Println(p2)

}
