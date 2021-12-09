package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		fmt.Println(".::desde teclado::.")
		countLines(os.Stdin, counts) // no indica archivo, leer teclado
	} else { // leer archivos
		fmt.Println(".::desde archivo::.")
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
			}
			countLines(f, counts)
			f.Close()
		}
	}
	// mostrar el resultado
	for line, n := range counts {
		fmt.Printf("%d: %s\n", n, line)
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
