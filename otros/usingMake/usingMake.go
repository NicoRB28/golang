package main

import (
	"fmt"
)

func main() {
	//usar make es mas una cuestion de eficiencia, al usarlo puedo delimitar una capacidad de
	//inicio en este caso es 100 por ende el array que esta por debajo del slice tiene una
	//capacidad inicial de 100 por mas que el length del slice sea 10 esto hace que al superar los 10
	//el esfuerzo de aumentar el array disminuye dado que recien al superar 100 tiene que aumentar de tamaño
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println("tamaño: ", len(x))
	fmt.Println("capacidad:", cap(x))

	x = append(x, 34, 66)
	fmt.Println(x)
	fmt.Println("tamaño: ", len(x))
	fmt.Println("capacidad:", cap(x))

	//en la siguiente linea, se supera la capacidad del array de fondo por ende se debe ampliar
	//la capacidad de duplica de 12 pasa a 24
	x = append(x, 99)
	fmt.Println(x)
	fmt.Println("tamaño: ", len(x))
	fmt.Println("capacidad:", cap(x))

}
