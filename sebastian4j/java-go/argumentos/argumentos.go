package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("una forma de for\n")
	for index, arg := range os.Args {
		if index == 0 {
			fmt.Printf("Program: %s\n", arg)
		} else {
			fmt.Printf("Argument: %d: %s\n", index, arg)
		}
	}
	fmt.Printf("otra forma de for\n")
	for index := 0; index < len(os.Args); index++ {
		fmt.Printf("index: %d value: %s\n", index, os.Args[index])
	}
	fmt.Printf("for y switch 1\n")

	for index, arg := range os.Args {
		switch {
		case index == 0:
			fmt.Printf("Program: %s\n", arg)
		default:
			fmt.Printf("Argument: %d: %s\n", index, arg)
		}
	}

	fmt.Printf("for y switch 2\n")
	for index, arg := range os.Args {
		switch index {
		case 0:
			fmt.Printf("Program: %s\n", arg)
		default:
			fmt.Printf("Argument: %d: %s\n", index, arg)
		}
	}
}
