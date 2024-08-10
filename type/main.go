package main

import "fmt"

type TestInterface interface {
	Hello()
}

type TestStruct struct {
	Name string
}

func (t *TestStruct) Hello() {
	fmt.Println("Hello", t.Name)
}

type TestStruct2 struct {
	age int
}

func (t *TestStruct2) Hello() {
	fmt.Println("Hello ", t.age)
}

func main() {

	var t TestInterface

	var testStruct *TestStruct

	t = testStruct
	t.Hello()

	t = &TestStruct{Name: "Tri"}
	t.Hello()

	t = &TestStruct2{age: 23}
	t.Hello()
}
