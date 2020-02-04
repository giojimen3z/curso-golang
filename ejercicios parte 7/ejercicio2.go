package main

import (
	"fmt"
)

type persona struct {
	nombre string
}

func main() {

	p1 := persona{
		nombre: "gio",
	}

	fmt.Println(p1)

	cambiame(&p1)
	fmt.Println(p1)

}

func cambiame(p *persona) {

	p.nombre = "johan"

}
