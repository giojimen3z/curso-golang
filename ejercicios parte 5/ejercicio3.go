package main

import (
	"fmt"
)

type vehiculo struct {
	puertas int
	color   string
}

type camion struct {
	vehiculo
	cuatroRuedas bool
}

type sedan struct {
	vehiculo
	lujoso bool
}

func main() {

	cam := camion{

		vehiculo: vehiculo{

			puertas: 2,
			color:   "Rojo",
		},
		cuatroRuedas: true,
	}

	sed := sedan{

		vehiculo: vehiculo{

			puertas: 4,
			color:   "Negro",
		},
		lujoso: true,
	}

	fmt.Println(cam)
	fmt.Print(sed)

}
