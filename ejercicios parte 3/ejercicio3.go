package main

import (
	"fmt"
)

func main() {

	na := 1993
	age := 0
	for na < 2019 {
		age++
		fmt.Println(na)
		na++

	}

	fmt.Println("Su edad es: ", age)
}
