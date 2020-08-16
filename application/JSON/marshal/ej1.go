package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type ColorGroup struct {
	ID     int
	Name   string
	Colors []string
}

func main() {
	group := ColorGroup{
		ID:     1,
		Name:   "Favoritos",
		Colors: []string{"azul", "verde", "rojo"},
	}

	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println(err)
	}
	os.Stdout.Write(b)

}
