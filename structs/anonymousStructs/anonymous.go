package main

import (
	"fmt"
)

func main() {

	//no creo un type person struct{}, directamente creo el struct sin un nuevo tipo
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1.first, " ", p1.last, " ", p1.age)

}
