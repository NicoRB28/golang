package main

import (
	"fmt"
)

func main() {

	f := func() {
		fmt.Println("mi first func expression")
	}

	fp := func(num int) {
		fmt.Println("second func expression with param ", num)
	}

	f()

	fp(28)

}
