package main

import "fmt"

type person struct {
	first       string
	last        string
	favIceCream []string
}

func main() {

	p1 := person{
		first: "Nicolas",
		last:  "Boschi",
		favIceCream: []string{
			"chocolate",
			"dulce de leche",
			"limon a la crema",
		},
	}
	p2 := person{
		first: "Pepe",
		last:  "Argento",
		favIceCream: []string{
			"tramontana",
			"don pedro",
			"frutilla",
		},
	}

	var m map[string]person
	m = map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for key, value := range m {
		fmt.Println("apellido ", key, ":")
		fmt.Println("primer nombre: ", value.first, "\ngustos de helado: ")
		for i, iceCreame := range value.favIceCream {
			fmt.Println(i, ")", iceCreame)
		}
	}

}
