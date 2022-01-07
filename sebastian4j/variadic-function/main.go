package main

import "fmt"

func main() {
	fmt.Println(sum())
	fmt.Println(sum(3, 4, 5))
	fmt.Println(sum2([]int{3, 4, 5}))
	fmt.Printf("%T\n", sum)
	fmt.Printf("%T\n", sum2)
}

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func sum2(vals []int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}
