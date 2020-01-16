package main

import (
	"fmt"
)

func main() {

	switch {

	case false:
		fmt.Println("No deberia imprimir")
	case true:
		fmt.Println("Deberia imprimir")

	}
}
