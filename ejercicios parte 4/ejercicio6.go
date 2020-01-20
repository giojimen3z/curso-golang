package main

import (
	"fmt"
)

func main() {

	x := map[string][]string{

		`gio`:    []string{`lee`},
		`andres`: []string{`jugar`},
		`johan`:  []string{`salar`},
	}

	for i, val := range x {

		fmt.Println("Registro: ", i)

		for j, reg := range val {

			fmt.Println(j, reg)

		}
	}

}
