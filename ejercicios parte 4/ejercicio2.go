package main

import (
	"fmt"
)

func main() {

	x := []int{1, 2, 3, 4, 5, 11, 25, 4}

	for i, v := range x {

		fmt.Println(i, v)
	}

	fmt.Printf("el tipo del arry es: %T", x)
}
