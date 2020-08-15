package main

import (
	"fmt"
)

func main() {

	slice1 := []string{"chocolate", "helado", "chocotorta"}
	slice2 := []string{"arroz", "pollo", "pizza"}

	fmt.Println(slice1)
	fmt.Println(slice2)

	sliceDeSlice := [][]string{slice1, slice2}
	//[[chocolate helado chocotorta]]
	//[[arroz     pollo  pizza]]
	fmt.Println(sliceDeSlice)

	fmt.Println(sliceDeSlice[0][1]) //imprime helado
	fmt.Println(sliceDeSlice[1][2]) //imprimir pizza

}
