package main

import (
	"encoding/json"
	"fmt"
)

//los atributos del struct tienen que empezar con mayuscula para que marshal funcione
type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{
		First: "James",
		Last:  "Bond",
		Age:   32,
	}
	p2 := person{
		First: "Pepe",
		Last:  "Argento",
		Age:   40,
	}

	people := []person{
		p1,
		p2,
	}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(bs))

}
