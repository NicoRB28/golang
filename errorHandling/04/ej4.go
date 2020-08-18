package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	f, err := os.Open("../03/names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer f.Close()

	bs, e := ioutil.ReadAll(f)
	if e != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs))
}
