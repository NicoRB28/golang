package main

import (
	"fmt"
	"sort"
)

func main() {
	slice1 := []int{7, 3123, 53, 231, 43, 5, 21, 65, 321}
	slice2 := []int{8421, 321, 45, 23, 53, 1, 354, 53}
	slice3 := []string{"hola", "arbol", "manzana", "miel", "lechuga"}

	sort.Ints(slice1)
	fmt.Println(slice1)

	sort.Slice(slice2, func(i, j int) bool { return slice2[i] > slice2[j] })
	fmt.Println(slice2)
	fmt.Println(slice3)
	sort.Strings(slice3)
	fmt.Println(slice3)
}
