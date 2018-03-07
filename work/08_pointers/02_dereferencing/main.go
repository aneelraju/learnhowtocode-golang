package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // pointer address

	var b *int = &a

	fmt.Println(b)  // pointer address; referencing
	fmt.Println(*b) // 43; de-referencing

	// b is an int pointer
	// b points to the memory address where an int is stored
	// to see the value in that memory address, add a * in front of b
	// this is know as dereferncing
	// the * is an operator in this case
}
