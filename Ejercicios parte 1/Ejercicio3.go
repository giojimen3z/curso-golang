package main

import (
	"fmt"
)
var a int
var b string
var c bool

func main() {
	
	a = 42
	b = "James Bond"
	c = true
	fmt.Printf("%v \n",a)
	fmt.Printf("%v \n",b)
	fmt.Printf("%v \n",c)
	s1 := fmt.Sprintf(" %v\t %v\t %v\t ",a, b, c)
	fmt.Println(s1)
}
