package main

import (
	"fmt"
)

func main() {

	personas := map[string][]string{
		"Boschi_Nico":  []string{"programar", "tocar guitarra", "estudiar"},
		"Argento_Pepe": []string{"futbol", "comer", "dormir"},
		"Alguien_mas":  []string{"no se", "que", "poner"},
	}

	personas["una_persona_nueva"] = []string{"algo", "mas", "para agregar"}

	for key, value := range personas {
		fmt.Println(key, ":")
		for index, v := range value {
			fmt.Println("index: ", index, " gusto: ", v)
		}
	}

	delete(personas, "Boschi_Nico")

	fmt.Println("Despues de borrar:")

	for key, value := range personas {
		fmt.Println(key, ": ")
		for index, v := range value {
			fmt.Println("index: ", index, " gusto: ", v)
		}
	}

}
