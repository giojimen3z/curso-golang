package main

import (
	"fmt"
)

func main() {

	defer foo()

	fmt.Println("Hello, playground")
}

func foo() {

	defer func() {

		fmt.Println("foo diferida  funciono")

	}()

	fmt.Println("Foo funciono")

}
