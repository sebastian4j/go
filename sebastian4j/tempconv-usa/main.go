package main

import (
	"fmt"
	"os"

	"github.com/sebastian4j/tempconv"
)

// las posiciones se inicializan en 0
var pc [256]byte

/*
go mod edit -replace github.com/sebastian4j/tempconv=../tempconv
go mod tidy
*/
func main() {
	valor := tempconv.CToF(67)
	fmt.Println(valor)
}

func init() {
	fmt.Println("init 1")
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&2)
		fmt.Printf("después: %d: %v\n", i, pc[i])
	}
}

func init() {
	fmt.Println("init 2")
	cwd, err := os.Getwd()
	fmt.Printf("%v %v\n", cwd, err)

	if cwd, err := os.Getwd(); err == nil {
		fmt.Println(cwd)
	}
}

/*
 scope:
el alcance de una declaración es una region de texto de programa, es una propiedad de tiempo de compilación, un nombre declarado en un bloque no es visible fuera de el. el bloque encierra la declaración y determina su alcance.

lifetime:
- es el rango de tiempo durante la ejecución en el que la variable puede ser referenciada por otras partes del programa (runtime property)

universe block: todo el source code
*/
