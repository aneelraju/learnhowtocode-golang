package main

import (
	"fmt"
)

func zero(x int) {
	fmt.Printf("%p\n", &x) // address in func zero
	fmt.Println(&x)        // address in func zero
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)
	// address in main;
	// %p pointer base 16 notation, with leading 0x
	fmt.Println(&x) // address in main
	zero(x)
	fmt.Println(x) // x is still 5
}
