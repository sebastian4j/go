package main

import (
	"flag"
	"fmt"
	"strings"
)

type myFlag struct {
	valor string
}

func (mf *myFlag) Set(s string) error {
	fmt.Printf("nuevo valor seteado en flag: %s\n", s)
	mf.valor = s
	return nil
}

func (mf *myFlag) String() string {
	return mf.valor
}

// ./flags -s Xs -n e r t y
// eXsrXstXsy

// para usar el nuevo flag agregado posteriormente:
// go run . -s Xs -fg fg-val   a s

// sep y n son punteros a las variables flag
// acceder indirectamente a sus valores con *n y *sep
// el booleano no espera un valor, si esta presente es true
var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {
	// agregar un flag extra
	mf := myFlag{"default ng"}
	flag.CommandLine.Var(&mf, "fg", "usage mf") // nombre + uso

	//////////////

	flag.Parse() // tiene que ser llamado antes de ser usado
	fmt.Printf("n: %v sep: %v nuevo: %v\n", *n, *sep, mf.valor)
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
	for _, j := range flag.CommandLine.Args() {
		fmt.Println(j)
	}
}
