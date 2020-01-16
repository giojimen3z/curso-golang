package main

import (
	"fmt"
)

func main() {

	x := "Robin"
	if x == "batman" {

		fmt.Printf("el valor de la variable x es: %v", x)
	} else if x == "Robin" {
		fmt.Printf("el valor de la variable x es: %v", x)
	} else {
		fmt.Println("Ninguno")
	}
}
