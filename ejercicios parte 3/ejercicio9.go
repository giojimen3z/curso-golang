package main

import (
	"fmt"
)

func main() {

	deportefav := "futbol"

	switch deportefav {

	case "futbol":
		fmt.Println("ve al estadio")
	case "beisbol":
		fmt.Println("ve a la playa")

	case "boleiball":
		fmt.Println("quedate en casa")
	default:
		fmt.Println("que te gusta hacer?")

	}
}
