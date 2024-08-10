package main

import (
	"fmt"
)

type Man struct {
	name string
	age  int
}

func testStructType(man *Man) {
	fmt.Println(man.age)
	fmt.Println(man.name)
}

func changeData(x *int) {
	a := x

	*a = 100
}

func mapDelta(md map[string]int, s string) {
	md[s] = 33
}

func main() {
	b := 255
	a := &b

	fmt.Println("a", a)
	fmt.Println("b", *a)

	// Dereferencing
	fmt.Println("Dereferencing", *&*&b)

	tri := &Man{"tri", 23}

	var man *Man

	man = tri

	fmt.Println("Tri", tri)
	fmt.Println("Man", man.age)

	testStructType(tri)

	changeData(&b)
	fmt.Println(b)

	// Map
	m := make(map[string]int)

	m["Tristan"] = 42
	fmt.Println(m["Tristan"])
	mapDelta(m, "Tristan")
	fmt.Println(m["Tristan"])
}
