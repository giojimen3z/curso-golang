package Ejercicios Parte 9

import(

	"fmt"
)

func main()  {
	c := make(chan int)

	c <- 42

	fmt.Println(<-c)
}

