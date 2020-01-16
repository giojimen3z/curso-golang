package main

import (
	"fmt"
)

func main() {

	s1 := `Esto es un 
	string  
	literal 
	no interpretado`
	s2 := "esto e sun string literal interpretado"

	fmt.Println(s1)
	fmt.Println(s2)

}
