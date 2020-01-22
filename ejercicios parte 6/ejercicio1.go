package main

import (
	"fmt"
)

func main() {

	n := foo()
	x, s := bar()
	fmt.Println(n)
	fmt.Println(s, x)

}

func foo() int {

	return 42

}

func bar() (int, string) {

	return 1942, "en esta fecha se descubrio america"

}
