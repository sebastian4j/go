package main

import (
	"fmt"
	"sync"
)

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func main() {
	//var x map[string]string // no se puede acceder a un elemento, no está inicializado
	//var y []string // queda en largo 0
	var z [2]string // queda inicializado
	//fmt.Println(x["s"])
	//fmt.Println(y[0])
	fmt.Println(z[0])

	cache.Lock()
	cache.mapping["a"] = "aa"
	v := cache.mapping["a"]
	cache.Unlock()
	fmt.Println(v)
}
