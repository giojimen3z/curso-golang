package main

import (
	"fmt"
)

func main() {

	for i := 10; i < 100; i++ {

		fmt.Printf("Cuando dividimos %v entre 4 el modulo  es: %v\n", i, i%4)
	}
}
