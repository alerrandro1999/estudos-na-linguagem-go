package main

import (
	"fmt"
)

var t = "texto"

func maina() {
	x := 10.033
	y := "ola bom dia"

	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)

	// x = 20
	// fmt.Printf("x: %v, %T\n", x, x)

	x, z := 20, 30
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("z: %v, %T\n", z, z)
}
