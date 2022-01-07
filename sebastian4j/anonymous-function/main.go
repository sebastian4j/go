package main

import (
	"fmt"
	"strings"
)

func main() {
	var suma int32 = 1
	fmt.Println(strings.Map(
		func(r rune) rune { return r + suma }, "hola"))

	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func squares() func() int {
	var x int
	// closures
	return func() int {
		x++
		return x * x
	}
}
