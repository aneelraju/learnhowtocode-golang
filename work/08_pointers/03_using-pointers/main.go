package main

import "fmt"

func main() {

	a := 43

	fmt.Println(a)  // 43
	fmt.Println(&a) // pointer

	var b *int = &a

	fmt.Println(b)  // pointer
	fmt.Println(*b) // 43

	*b = 42        // pass by value
	fmt.Println(a) // 42

	// this is useful
	// we can pass a memory address instead of a bunch of values
	// and then we can still change the value of whatever is stored
	// at that memory address
	// this makes our programs more performant
	// we don't have to pass around large amounts of data
}
