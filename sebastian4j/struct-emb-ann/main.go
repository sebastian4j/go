package main

import "fmt"

type Point struct {
	X, Y int
}

type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}

func main() {
	var w1 Wheel
	w1.X = 8 // es anonimo (nombre implicito, el tipo), pasa directo
	w1.Y = 9
	fmt.Println(w1)
	// pero tambien permite acceder usando el nombre del tipo
	w1.Circle.Point.X = 88
	w1.Circle.Point.Y = 99

	fmt.Printf("%#v\n", w1)

	fmt.Println(Wheel{Circle{Point{1, 2}, 3}, 4})
	fmt.Println(Wheel{
		Circle: Circle{
			Point:  Point{1, 2},
			Radius: 5, // la coma es necesaria
		},
		Spokes: 6, // la coma es necesaria
	})
}
