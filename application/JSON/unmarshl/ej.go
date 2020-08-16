package main

import (
	"encoding/json"
	"fmt"
)

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
	fmt.Println("transformado en json")
	fmt.Println(string(bs))

	var people2 []person

	err2 := json.Unmarshal(bs, &people2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println("devuelvo al estado original: (unmarshal)")
	fmt.Println(people2)

}
