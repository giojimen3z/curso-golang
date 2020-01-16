package main

import (
	"fmt"
)

func main() {

	a := (42 == 42)
	b := (42 >= 40)
	c := (42 <= 45)
	d := (42 != 45)
	e := (42 < 889)
	f := (42 > 29)

	fmt.Println(a, b, c, d, e, f)
}
