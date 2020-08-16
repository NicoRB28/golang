package main

import "fmt"

func main() {

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := evenSum(sum, numbers...)
	fmt.Println("suma de pares:", result)

	multResult := evenSum(mult, numbers...)
	fmt.Println("multiplicacion de pares: ", multResult)

	sumaDeTodos := sum(numbers...)
	fmt.Println("suma de todos los valores:", sumaDeTodos)
}

func sum(x ...int) int {
	tot := 0
	for _, value := range x {
		tot += value
	}
	return tot
}

func mult(x ...int) int {
	tot := 1
	for _, value := range x {
		tot *= value
	}
	return tot
}

func evenSum(f func(x ...int) int, nums ...int) int {
	var slice []int
	for _, value := range nums {
		if value%2 == 0 {
			slice = append(slice, value)
		}
	}

	return f(slice...)
}
