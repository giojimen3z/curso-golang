package main

import (
	"fmt"
)

func main() {

	x := make([]string, 50, 50)
	x = []string{`a`, `b`, `c`, `d`, `e`, `f`, `g`, `h`, `i`, `j`, `k`, `l`, `m`, `n`, `ñ`, `o`, `p`, `q`, `r`, `s`, `t`, `u`, `v`}

	fmt.Println(len(x))
	fmt.Print(cap(x))

	for i := 0; i < len(x); i++ {

		fmt.Println(i, x[i])

	}

}
