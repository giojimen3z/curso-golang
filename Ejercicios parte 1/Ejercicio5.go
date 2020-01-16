package main

import (
	"fmt"
)

type gio int

var x gio

var y int

func main() {

	fmt.Println(x)
	fmt.Printf("El tipo de x es: %T \n", x)
	x = 42
	fmt.Println(x)

	y = int(x)

	fmt.Println("EL valor de Y es:", y)
	fmt.Printf("El tipo de Y es: %T\n", y)
}
