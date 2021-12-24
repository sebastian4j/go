package main

import (
	"fmt"
	"strings"
)

func main() {
	months := [...]string{1: "Enero", 2: "Febrero", 3: "Marzo", 4: "Abril", 5: "Mayo"}
	m1 := months[1:3]
	fmt.Printf("%v [0]: %s [1:3]:%v\n", months, months[0], m1)
	months[1] = strings.ToUpper(months[1])
	fmt.Printf("%v [0]: %s [1:3]:%v\n", months, months[0], m1)

	fmt.Println(
		cap(m1), // la capacidad del array al que referencia
		len(m1)) // el largo del arreglo que tiene la referencia

	m2 := months[3:]

	fmt.Println(m2)

	a := [...]int{0, 1, 2, 3, 4, 5} // array
	b := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a == b) // son comparables
	fmt.Printf("%T %T\n", a[:], a)
	reverse(a[:])
	fmt.Println(a)
	s := []int{0, 1, 2, 3, 4, 5} // no comparable
	s = append(s, 6)
	reverse(s)
	fmt.Println(s)
	fmt.Println(cap(s), len(s))

	fmt.Println(cap(s), len(s))

	var ss []int
	fmt.Println(cap(ss), len(ss))

	rn := []rune("hello")
	fmt.Printf("%c \n", rn)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%2d cap=%2d\t%v\n", i, cap(y), y)
		x = y
	}
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func appendInt(x []int, y ...int) []int {
	var z []int
	// zlen := len(x) + 1 // cuando y es int
	zlen := len(x) + len(y)
	if zlen < cap(x) { // tiene capacidad
		z = x[:zlen] // se obtiene un slice con la capacidad necesaria
	} else {
		// no queda espacio, ubicar un nuevo arreglo
		// crece el doble, amortiza complejidad lineal
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	// agregar el nuevo valor
	// z[len(x)] = y // cuando y es int
	copy(z[len(x):], y) // y es ... int
	return z
}
