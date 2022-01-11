package main

import (
	"fmt"
)

/*

https://gist.github.com/asukakenji/f15ba7e588ac42795f421b48b8aede63


export GOOS=windows
export GOARCH=amd64

*/

type ByteCounter int

var funciones = make(map[string]func())

func init() {
	funciones["metodo"] = metodo
}

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

func main() {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
	invoca(funciones["metodo"])
}

func invoca(fn func()) {
	fn()
}

func metodo() {
	fmt.Println("método invocado")
}
