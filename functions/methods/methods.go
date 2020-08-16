package main

import "fmt"

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

//antes de los parametros se pone quien recibe linkea el speak al typo del receiver
func (s secretAgent) speak() {
	//al estar linkeados el metodo puede acceder a las variables del struct
	fmt.Println("I am", s.first, s.last)
}

func main() {
	sa1 := secretAgent{
		person: person{
			"james",
			"bond",
		},
		ltk: true,
	}
	fmt.Println(sa1)
	sa1.speak()

	sa2 := secretAgent{
		person: person{
			"Pepe",
			"Argento",
		},
		ltk: true,
	}

	sa2.speak()
}
