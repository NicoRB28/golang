package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	var p4, p1, p2, p3 person
	constructor(&p1, "Julio", "Profe", 50)
	constructor(&p2, "Christian", "Rodriguez", 45)
	constructor(&p3, "Maria", "Sola", 48)
	constructor(&p4, "Fernando", "Carlos", 40)

	people := []person{
		p1,
		p2,
		p3,
		p4,
	}

	err := json.NewEncoder(os.Stdout).Encode(people)

	if err != nil {
		fmt.Println(err)

	}

}

func constructor(va *person, f string, l string, a int) {
	va.First = f
	va.Last = l
	va.Age = a
}
