package main

import "fmt"

func main() {

	fmt.Println(recursiveFactorial(4))
	fmt.Println(loopFactorial(4))

	fmt.Println(recursiveFactorial(6))
	fmt.Println(loopFactorial(6))
}

func loopFactorial(num int) int {
	resultado := 1
	for i := num; i > 0; i-- {
		resultado = resultado * i
	}
	return resultado
}
func recursiveFactorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * recursiveFactorial(num-1)
}
