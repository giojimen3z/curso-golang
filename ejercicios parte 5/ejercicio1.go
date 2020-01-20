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

	p := persona{

		nombre:     "gio",
		apellido:   "jimenez",
		saboresfav: []string{"verde", "blanco", "rojo"},
	}

	fmt.Println(p)
}
