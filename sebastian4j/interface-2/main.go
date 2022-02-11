package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = true

func main() {
	// es un puntero su tipo no es nul, es un *bytes.Buffer
	// pero su valor es null (apunta a nada)
	var buf *bytes.Buffer
	if debug {
		buf = new(bytes.Buffer) // crea un puntero
	}
	fmt.Println("1")
	f(buf)

	var fub bytes.Buffer // empty buffer, listo para usar
	fmt.Println("2")
	f(&fub)

	var bfu io.Writer
	fmt.Println("3")
	f(bfu)
}

func f(out io.Writer) {
	if out != nil {
		fmt.Println("no nulo")
		out.Write([]byte("hola"))
	} else {
		fmt.Println("nulo")
	}
}
