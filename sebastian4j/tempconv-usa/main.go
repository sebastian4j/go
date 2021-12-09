package main

import (
	"fmt"

	"github.com/sebastian4j/tempconv"
)

/*
go mod edit -replace github.com/sebastian4j/tempconv=../tempconv
go mod tidy
*/
func main() {
	valor := tempconv.CToF(67)
	fmt.Println(valor)
}
