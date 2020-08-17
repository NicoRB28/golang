package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Lucho",
		Age:   59,
	}

	u2 := user{
		First: "Roberto",
		Age:   43,
	}

	u3 := user{
		First: "Jose",
		Age:   32,
	}

	users := []user{
		u1,
		u2,
		u3,
	}

	fmt.Println(users)

	js, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(js))

	//para hacer el unmarsh los pasos son:
	//1.crear una variable para pasar un puntero a ella en la funcion unmarsh, ahi seran almacenados los elementos
	var usersUnmarshal []user
	//llamar a json.unmarsh(JSON, laDireccionDeLaVariable)
	unmarError := json.Unmarshal(js, &usersUnmarshal)
	if unmarError != nil {
		fmt.Println(unmarError)
		return
	}

	fmt.Println(usersUnmarshal)

}
