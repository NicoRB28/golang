package main

import "fmt"

func main() {

	fmt.Println("2+3= ", mySum(2, 3))
	fmt.Println("4+7= ", mySum(4, 7))
	fmt.Println("8+5 = ", mySum(8, 5))

}

func mySum(x ...int) int {

	tot := 0
	for _, value := range x {
		tot += value
	}

	return tot
}
