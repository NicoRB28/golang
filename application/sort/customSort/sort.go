package main

import (
	"fmt"
	"sort"
)

type person struct {
	first string
	age   int
}

type byAge []person

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].age < a[j].age }

type byName []person

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].first > a[j].first }

func main() {
	p1 := person{
		first: "Nicolas",
		age:   28,
	}
	p2 := person{
		first: "Fernando",
		age:   35,
	}
	p3 := person{
		first: "Roberto",
		age:   80,
	}

	people := []person{
		p3,
		p2,
		p1,
	}

	fmt.Println(people)

	sort.Slice(people, func(i, j int) bool { return people[i].age < people[j].age })
	fmt.Println(people)

	p4 := person{
		first: "Matias",
		age:   43,
	}
	p5 := person{
		first: "Fausto",
		age:   35,
	}
	p6 := person{
		first: "Pepe",
		age:   80,
	}

	people2 := []person{
		p6,
		p4,
		p5,
	}

	fmt.Println("antes del sort:", people2)
	sort.Sort(byAge(people2))
	fmt.Println("despues del sort: ", people2)

	sort.Sort(byName(people2))
	fmt.Println(people2)

	//sort.Sort(data Interface)--> Sort recibe un dato de tipo Interface, este tipo de dato implementa los
	//3 metodos que estan desarrollados mas arriba y que posibilitan al Sort ordenarlos
}
