package main

import "fmt"

func main() {

	var x int
	fmt.Println(x)
	x++
	fmt.Println(x)

	{
		//esta declaracion es valida dado que esta en un bloque interno
		var x int = 100
		fmt.Println(x)
		var y int = 200
		fmt.Println(y)
	}
	//fmt.Println(y) ---> error porque en este scope y no existe

	fmt.Println(x)

}
