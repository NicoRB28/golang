package main

import (
	"fmt"
)

func main() {

	slice := []int{}
	slice2 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)

	slice = append(slice, slice2...)
	fmt.Println(slice)

	slice = append(slice[:1], slice[4:]...)
	//elimina el 2,3 y 4
	//el comando [a:b] toma desde a (inclusive) hasta b(no inclusive)
	fmt.Println(slice)

}
