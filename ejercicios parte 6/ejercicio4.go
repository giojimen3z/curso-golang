package main

import (
	"fmt"
)

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func (p persona) presentar() {

	fmt.Println("Hola Soy: ", p.nombre, p.apellido, "mi edad es:", p.edad, "años")
}

func main() {

	p1 := persona{

		nombre:   "james",
		apellido: "bond",
		edad:     32,
	}

	p1.presentar()

}
