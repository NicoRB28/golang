package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}

	p2 := person{
		first: "Otra",
		last:  "persona",
		age:   20,
	}
	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println("primer nombre: ", p1.first, " apellido: ", p1.last, " edad:  ", p1.age, " licencia para arma: ", p1.ltk)
}
