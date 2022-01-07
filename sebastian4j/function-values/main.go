package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Map(add, "hola"))
}

func add(r rune) rune {
	return r + 1
}
