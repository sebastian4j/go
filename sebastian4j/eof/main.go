package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			log.Println("eof alcanzado") // ctrl + d
			break
		}
		if err != nil {
			log.Fatalf("error al leer: %v", err)
		}
		log.Printf("%*s", 4, string(r))
	}
}
