package main

import (
	"fmt"
	"math"
)

type square struct {
	b int
	h int
}

type circle struct {
	r float64
}

type shape interface {
	area() float64
}

func (s square) area() float64 {
	return float64(s.b * s.h)
}
func (c circle) area() float64 {
	return (math.Pi * float64(c.r*c.r))
}

func info(s shape) {
	fmt.Println("mi area es: ", s.area())
}

func main() {
	sq := square{
		b: 2,
		h: 2,
	}

	ci := circle{
		r: 5,
	}

	info(sq)
	info(ci)
}
