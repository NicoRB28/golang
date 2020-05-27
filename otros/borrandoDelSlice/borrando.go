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
	fmt.Println(slice)

}
