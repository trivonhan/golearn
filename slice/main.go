package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 40, 50}

	fmt.Println(s)

	for i, a := range s {
		fmt.Printf("Index: %d, Value: %d\n", i, a)
	}
}
