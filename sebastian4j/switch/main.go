package main

import (
	"fmt"
)

func main() {
	switch coinflip() {
	case "heads":
		fmt.Println("heads")
	case "tails":
		fmt.Println("tails")
		fmt.Println("sliat")
	default:
		fmt.Println("default")
	}

	x := 0
	// tagless switch
	switch {
	case x > 0:
		fmt.Println(">0")
	default:
		fmt.Println("= 0")
	case x < 0:
		fmt.Println("< 0")
	}
}

func coinflip() string {
	return "tails"
}
