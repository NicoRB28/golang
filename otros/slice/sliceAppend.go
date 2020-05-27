package main

import (
	"fmt"
)

func main() {

	slice := []int{}
	slice2 := []int{9, 10, 11, 12}
	slice = append(slice, 1, 2, 3, 4, 5, 6, 7, 8)

	fmt.Println(slice)

	slice = append(slice, slice2...)
	fmt.Println(slice)
}
