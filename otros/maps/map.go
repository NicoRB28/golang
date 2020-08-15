package main

import (
	"fmt"
)

func main() {

	m := map[int]string{
		1: "Nico",
		2: "Fernando",
		3: "Pepe",
	}

	fmt.Println(m)

	fmt.Println(m[1])
	//fmt.Println(m[5]) //por default si la key no esta en el map imprime el valor vacio correspondiente al tipo
	//para evitar esto se puede hacer lo siguiente:

	//value, ok := m[5]
	//if ok {
	//	fmt.Println(value)
	//} else {
	//	fmt.Println("no se encuentra la key en el map")
	//}
	//eso se puede reEscribir asi:
	if value, ok := m[5]; ok {
		fmt.Println(value)
	} else {
		fmt.Println("no esta la key en el map")
	}
	fmt.Println(m[2])

	fmt.Println("imprimiendo el map con un for: ")
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

	//para agregar un nuevo elmento:
	m[5] = "Roberto"
	fmt.Println("despues de agregar un elemento:")
	for key, value := range m {
		fmt.Println(key, ":", value)
	}

}
