package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func main() {
	//la idea es que si tengo que procesar muchos datos poder procesarlos a la par en
	//procesos separados y despues juntar los resultados

	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fanOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("end")

}

func populate(c1 chan int) {
	for i := 0; i < 100; i++ {
		c1 <- i
	}
	close(c1)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const goroutines = 10
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		fmt.Println("goroutines: ", runtime.NumGoroutine())
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- timeConsumingWork(v2)
				}(v)
			}
			wg.Done()
		}()

	}
	wg.Wait()
	close(c2)
}

func timeConsumingWork(v int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return v + rand.Intn(1000)
}
