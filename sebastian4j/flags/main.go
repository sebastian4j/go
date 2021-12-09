package main

import (
	"flag"
	"fmt"
	"strings"
)

// ./flags -s Xs -n e r t y
// eXsrXstXsy

// sep y n son punteros a las variables flag
// acceder indirectamente a sus valores con *n y *sep
// el booleano no espera un valor, si esta presente es true
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	flag.Parse() // tiene que ser llamado antes de ser usado
	fmt.Printf("n: %v sep: %v\n", *n, *sep)
	// flag.Args() son los argumentos no flags
	fmt.Print(strings.Join(flag.Args(), *sep))
	fmt.Println()

	// si esta presente el -n se ponen 3 enter
	if *n {
		fmt.Println()
		fmt.Println()
		fmt.Println()
	} else { // si no 1
		fmt.Println()
	}
}
