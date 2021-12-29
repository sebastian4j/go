package main

import "fmt"

type Employee struct {
	ID           int
	Name, Addess string
}

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// return &tree{value: value}
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}
func main() {
	var a Employee
	a.Name = "nombre a"
	fmt.Println(a)
	nombre := &a.Name
	*nombre = "nombre b"
	fmt.Println(a)

	var aa *Employee = &a
	aa.ID = 1
	fmt.Println(aa)
	modifica(&a)
	fmt.Println(a)

	fmt.Println(".::tree")
	var t = add(nil, 1)
	t = add(t, 2)
	t = add(t, 0)
	fmt.Println(t)

	fmt.Println(Employee{ID: 1})
}

func modifica(aaa *Employee) {
	aaa.Name = "nombre c"
}
