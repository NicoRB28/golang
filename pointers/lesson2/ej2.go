package main

import "fmt"

func main() {
	//como Go hace pasaje por valor si quiero modificar el valor de un parametro tengo que usar punteros

	x := 0
	foo(x)
	fmt.Println(x) //==> print 0

	bar(&x)        //==> paso la referencia
	fmt.Println(x) //==> print 500
}

func foo(y int) {
	fmt.Println(y) //==> print 0
	y = 43
	fmt.Println(y) //==> print 43
}

func bar(y *int) {
	//==> aca espero de parametro una referencia
	fmt.Println(*y) // ==> print 0
	*y = 500        // ==> aca asigno el valor 500 en la celda referenciada
	fmt.Println(*y) // ==> print 500

}
