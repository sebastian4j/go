package main

import (
	"fmt"
	"os"
)

// - named type: define un nuevo tipo aunque sea ocupen el mismo tipo inicial
// pero no son considerados el mismo tipo,
// - no se pueden comparar o combinar !!! (son tipos distintos)
// en expresiones aritmeticas
// - si ambos tienen el mismo tipo inicial es posible hacer una conversión
// - con los tipos definidos iniciales tambien es posible hacer una conversión (float a int, string a []byte ...)
type Celsius float64
type Fahrenheit float64

// type's methods:
// asocia al tipo Celcius el método llamado String
func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

// creacion de variables
var a, j = 1, "hola" // asignación

func main() {
	celsius := FToC(85.0)
	fmt.Println(celsius.String())
	var names []string
	var f, err = os.Open(".")
	f.Close()

	// short variable declaration:
	d, c := 0, 1 // declaración
	// tiene que decalara al menos una nueva variable
	// d o c ya podrían existir pero no ambas, una tiene que ser nueva

	t := 0
	fmt.Printf("%v %v %v - %v %v - %v\n", a, b, j, f, err, t)
	fmt.Printf("%v\n", names)
	fmt.Printf("%v %v\n", d, c)

	d, c, b = c, d, a // intercambiar los valores

	//
	var e = 1 // named variable
	var g *bool
	// *g = false // no se puede, actualmente g apunta a nil
	g = boleana() // ahora g apunta a algo
	*g = false    // se puede, ahora g apunta a algo
	fmt.Printf("%v %v\n", e, g)

	medals := []string{"gold", "silver", "bronze"}
	medals[0] = "cobre"
	fmt.Println(medals)

	// nil puede ser una referencia o interface
	// en una comparación el primero tiene que ser asignable al tipo del segundo o vice versa
}

// es inicializada antes del llamar al main
var b = 3

func boleana() *bool {
	var booleana bool = false
	return &booleana
}
