package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			if len(line) > 0 {
				counts[line]++
			}
		}
	}
	for line, n := range counts {
		fmt.Printf("%d: %s\n", n, line)
	}
}
