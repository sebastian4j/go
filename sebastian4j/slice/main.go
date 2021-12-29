package main

import "fmt"

func main() {
	// un slice con largo y capacidad inicial 10
	var a = make([]int, 10)
	for i, _ := range a {
		a[i] = i
	}
	fmt.Println("a:", a)
	// un slice de capacidad inicial 0
	var b []int
	b = append(b, 10)
	b = append(b, 11)
	modifica(b)
	fmt.Println("b:", b)
	copy(a[8:], b)
	fmt.Println(a)
	// agrega b en a
	a = append(a, b...)
	fmt.Println("nuevo a:", a)

	push := append(a, 12)

	fmt.Println("push:", push, "a: ", a)    // push
	fmt.Println("top de a:", a[len(a)-1])   // top de a (obtener)
	fmt.Println("pop en a: ", a[:len(a)-1]) // quitar sin obtenerlo
}

func modifica(bb []int) {
	bb[0] = 24
	bb = append(bb, 2)
}
