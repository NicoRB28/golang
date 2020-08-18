package main

import "fmt"

func main() {
	var input1, input2, input3 string

	fmt.Println("Name:")
	_, err := fmt.Scan(&input1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Fav food:")
	_, err = fmt.Scan(&input2)
	if err != nil {
		panic(err)
	}

	fmt.Println("last name:")
	_, err = fmt.Scan(&input3)
	if err != nil {
		panic(err)
	}

	fmt.Println(input1, input2, input3)

}
