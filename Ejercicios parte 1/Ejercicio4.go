package main

import (
	"fmt"
)

type gio int

var x gio

func main() {

	fmt.Println(x)
	fmt.Printf("El tipo de x es: %T \n", x)
	x = 42
	fmt.Println(x)
	

}
