package main

import "fmt"

// maps no son comparables con == , solo son iguales si ambos son nil

func main() {
	ages := make(map[string]int)
	ages["a"] = 2
	ages["b"] = 2
	fmt.Println(ages)

	if _, ok := ages["c"]; !ok {
		fmt.Println("c no está presente")
	}
	modifica(ages)
	fmt.Println(ages)
}

func modifica(m map[string]int) {
	m["qwerty"] = 1
}
