package main

import (
	"fmt"
)

func main() {
	var x int
	y := &x // y es de tipo *int. apunta a x

	fmt.Printf(`el valor original: %v
el que apunta al original:%v
el valor al que apunta: %v
`, x, y, *y)

	*y = 1 // cambiar el valor al que apunta
	fmt.Printf("pre: %d\n", x)
	changer(&x)
	fmt.Printf("post: %d\n", x)

	var z *int
	fmt.Printf("el puntero x defecto es nil: %v\n", z)
	// si apunta a una variable es != nil
	// 2 punteros son iguales si son nil o apuntan a la misma variable

	// a y b estan iniciados x defecto a 0
	// así que ambos tienen un puntero distinto de nil
	var a, b int
	fmt.Println(
		&a == &a,  // es el mismo puntero, true el mismo puntero
		&a == &b,  // son distintos punteros, es false
		&a == nil) // no es nil, apunta a algo

	var c = nextChanger()
	fmt.Printf("puntero recibido: %v\n", c)
	fmt.Printf("valor del puntero recibido: %v\n", *c)

	// es false xq cada llamada está generando un nuevo puntero, aunque tengan el mismo valor
	fmt.Println(nextChanger() == nextChanger())
	fmt.Println(*nextChanger() == *nextChanger()) // el valor es true

	fmt.Println("new(T)")
	// new(T) : *T
	// crea una unnamed variable de tipo T
	d := new(int) // es un *int
	fmt.Println(d)
	*d = 2
	fmt.Println(*d)
}

func changer(x *int) {
	*x++
	fmt.Printf("cambiado en la función: %d\n", *x)
}

func nextChanger() *int {
	v := 11
	return &v
}

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	var x = &dummy
	return x
}

func newInt3() *int {
	var dummy int
	return &dummy
}
