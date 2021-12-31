package main

import "fmt"

func main() {
	fmt.Println(returns())
	fmt.Println(bareReturn())
	fmt.Println(returnss())
}

func returns() (int, int) {
	return 1, 2
}

func bareReturn() (x, y int) {
	x, y = 3, 4
	return
}

func returnss() (x, y int) {
	return 5, 6
}
