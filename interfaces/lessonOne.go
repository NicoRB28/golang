package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

//keyword identifier type
type secretAgent struct {
	person
	ltk bool
}

func (sa secretAgent) speak() {
	fmt.Println("I am", sa.first, sa.last, " agent speak")
}

func (pe person) speak() {
	fmt.Println("I am", pe.first, pe.last, " human speak")
}

func bar(h human) {
	fmt.Println("estoy en la funcion bar:", h)
}

//Keyword identifier type
//todo tipo que tenga implementado speak() es tmb un humano
type human interface {
	speak()
}

func main() {

	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	p1 := person{
		"Nico",
		"Bos",
	}

	fmt.Println(p1)

	sa1.speak()
	p1.speak()
	//polimorfismo, bar recibe un Human ambos son human porque implementan speak
	bar(sa1)
	bar(p1)
}
