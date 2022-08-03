package main

import (
	"fmt"
)

func main() {
	x := "ola tudo bem"
	y := "tudo bem"
	z := fmt.Sprint(x, y)
	fmt.Println(z)
}
