package main

import (
	"fmt"
)

type person struct {
	first            string
	last             string
	favoriteIceCream string
}

func main() {

	//declaro e inicio al mismo tiempo
	p1 := person{
		first:            "Nicolas",
		last:             "Boschi",
		favoriteIceCream: "Chocolate",
	}

	//para recordar que la declaracion y la inicializacion pueden ser por separado
	var p2 person

	p2 = person{
		first:            "Roberto",
		last:             "Carlos",
		favoriteIceCream: "limon",
	}

	fmt.Println(p1)
	fmt.Println(p2)

}
