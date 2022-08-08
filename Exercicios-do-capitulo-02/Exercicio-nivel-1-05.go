package main

import (
	"fmt"
)

type novotipo int

var y int

var x novotipo

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	y = 42
	fmt.Println(y)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
