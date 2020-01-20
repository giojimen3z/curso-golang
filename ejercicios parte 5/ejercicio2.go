package main

import (
	"fmt"
)

type persona struct {
	nombre     string
	apellido   string
	saboresfav []string
}

func main() {

	p1 := persona{

		nombre:     "gio",
		apellido:   "jimenez1",
		saboresfav: []string{"verde1", "blanco1", "rojo"},
	}
	p2 := persona{

		nombre:     "johan",
		apellido:   "jimenez",
		saboresfav: []string{"verde", "blanco", "rojo"},
	}

	m := map[string]persona{

		p1.apellido: p1,
		p2.apellido: p2,
	}

	for k, v := range m {

		fmt.Println(k)
		fmt.Println(v)

	}

}
