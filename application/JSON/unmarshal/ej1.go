package main

import (
	"encoding/json"
	"fmt"
)

type animal struct {
	Name  string
	Order string
}

func main() {
	var jsonS = []byte(`[
		{"Name": "Platypus", "Order": "Monotremata"},
		{"Name": "Quoll",    "Order": "Dasyuromorphia"}
	]`)

	var animals []animal
	//unmarshal recibe el json y la referencia al slice a donde se debe guardar el resultado(&animals)
	err := json.Unmarshal(jsonS, &animals)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%v", animals)
}
