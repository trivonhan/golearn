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

func main() {
	b := 255
	a := &b

	fmt.Println("a", a)
	fmt.Println("b", *a)

	tri := &Man{"tri", 23}

	var man *Man

	man = tri

	fmt.Println("Tri", tri)
	fmt.Println("Man", man.age)

	testStructType(tri)
}
