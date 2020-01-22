package main

import (
	"fmt"
)

func main() {

	s := struct {
		nombre  string
		amigos  map[string]int
		bebidas []string
	}{

		nombre: "gio",
		amigos: map[string]int{

			"johan":  123,
			"andres": 456,
		},
		bebidas: []string{

			"limonada",
			"coke",
		},
	}

	fmt.Println(s)

}
