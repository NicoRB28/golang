package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close() // ya de entrada lo cierro pero difiero el cierre al final

	r := strings.NewReader("wassup")

	io.Copy(f, r)
}
