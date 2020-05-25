package main

import "fmt"

func main() {

	res := 0

	fmt.Println(`la suma de los multiplos de 3 o 5 entre los 100 primeros naturales es:`)

	for i := 0; i <= 100; i++ {

		if i%3 == 0 || i%5 == 0 {
			res += i
		}
	}

	fmt.Println(res)

}
