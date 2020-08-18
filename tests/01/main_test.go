package main

import (
	"fmt"
	"testing"
)

//el test es buena practica que este en el mismo paquete
//el test se corre con go test y nada mas

func TestMySum(t *testing.T) {
	//test va con mayusculas

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 2}, 4},
		test{[]int{21, 21}, 42},
		test{[]int{3, 4, 5}, 12},
		test{[]int{-1, 0, 1}, 0},
	}

	var visual []string

	for _, v := range tests {
		x := mySum(v.data...)
		if x != v.answer {
			t.Error("Expected ", v.answer, "Got ", x)
			visual = append(visual, "x")
		} else {
			visual = append(visual, ".")
		}
	}

	fmt.Println(visual) //basado en Rspec de ruby se me ocurrio generar un slice con puntos correspondientes a
	//los tests que pasan y x correspondientes a los tests que no pasan una pavada basicamente jajajaj pero
	//suma

}
