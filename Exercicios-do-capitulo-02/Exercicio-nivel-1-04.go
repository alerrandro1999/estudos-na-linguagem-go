package main

import (
	"fmt"
)

type novotipo int

var x novotipo

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
