package main

import (
	"fmt"
	"sort"
)

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func main() {
	var names1 = []string{"z", "d", "c"}
	names1 = append(names1, "x")
	sort.Sort(StringSlice(names1))
	fmt.Println(names1)

	var names2 = []string{"z", "d", "c"}
	sort.Strings(names2)
	fmt.Println(names2)
}
