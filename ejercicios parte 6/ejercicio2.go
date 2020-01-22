package main

import (
	"fmt"
)

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	n := foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	n2 := bar(ii)
	fmt.Println(n)
	fmt.Println(n2)

}

func foo(xi ...int) int {

	var sum int

	for _, v := range xi {

		sum += v

	}

	return sum

}

func bar(xi []int) int {

	var sum int

	for _, v := range xi {

		sum += v

	}

	return sum

}
