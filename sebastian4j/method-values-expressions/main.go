package main

import (
	"fmt"
	"math"
	"time"
)

type Rocket struct {
}

func (r *Rocket) Launch() {

}

type Point struct {
	X, Y float64
}

func (p Point) Distance(t Point) float64 {
	return math.Hypot(t.X-p.X, t.Y-p.Y)
}

func main() {
	p := Point{1, 2}
	t := Point{4, 6}
	var origin Point
	disctanceFromP := p.Distance
	fmt.Println(disctanceFromP(origin))
	fmt.Println(disctanceFromP(t))

	var r Rocket
	r.Launch()
	fmt.Println(r)
	rr := new(Rocket)
	rr.Launch()
	fmt.Println(rr)

	time.AfterFunc(1*time.Second, rr.Launch)
	time.AfterFunc(1*time.Second, r.Launch)

	distace := (*Point).Distance
	fmt.Println(distace(&p, t))
}
