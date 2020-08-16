package main

import (
	"fmt"
)

func main() {

	//tambien era valido: provincias := make([]string,23,30)
	//provincias = []string{nombres de las provincias}
	provincias := []string{
		"Tierra del fuego",
		"Santa Cruz",
		"Chubut",
		"Rio Negro",
		"Neuquen",
		"Mendoza",
		"San Juan",
		"Catamarca",
		"Cordoba",
		"La Pampa",
		"Buenos Aires",
		"Santa Fe",
		"Corrientes",
		"Entre Rios",
		"San Luis",
		"Misiones",
		"La Rioja",
		"Salta",
		"Tucuman",
		"Chaco",
		"Santiago del Estero",
		"Jujuy",
		"Formosa",
	}

	fmt.Println("el tamaño de mi slice es: ", len(provincias))
	fmt.Println("la capacidad de mi slice es: ", cap(provincias))

	for index := 0; index < len(provincias); index++ {
		fmt.Println(index, ": ", provincias[index])
	}
}
