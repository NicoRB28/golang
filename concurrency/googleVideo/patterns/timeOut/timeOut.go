package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := boring("Joe")
	//time.After devuelve un canal en este caso despues de 5 segundos
	timeOut := time.After(5 * time.Second)
	for {
		select {
		case s := <-c:
			fmt.Println(s)
		case <-timeOut:
			fmt.Println("the end of time")
			return
		}
	}

}

func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}
