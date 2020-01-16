package main

import (
	"fmt"
)

func main() {

	na := 1993
	age := 0
	for {
		na++

		if na > 2019 {

			break
		}
		fmt.Println(na)
		age++
	}

	fmt.Println("Su edad es: ", age)
}
