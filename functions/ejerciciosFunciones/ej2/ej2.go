package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println("el resultado de foo es:", foo(slice...))
	fmt.Println("el resultado de bar es:", bar(slice))
}

func foo(x ...int) int {
	tot := 0
	for _, value := range x {
		tot += value
	}
	return tot
}

func bar(x []int) int {
	tot := 0
	//tambien es valido aca usar range pero para variar se hizo de la forma mas "normal"
	for i := 0; i < len(x); i++ {
		tot += x[i]
	}
	return tot
}
