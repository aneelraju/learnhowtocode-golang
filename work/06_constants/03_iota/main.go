package main

import "fmt"

const (
	// iota increments on declaration
	A = iota // 0
	B = iota // 1
	C = iota // 2
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
}
