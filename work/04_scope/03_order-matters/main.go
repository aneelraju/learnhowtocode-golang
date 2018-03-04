package main

import "fmt"

func main() {
	fmt.Println(x)
	fmt.Println(y)
	x := 42 // order matters
}

var y = 42 // order doesn't matter; outerscope variables are accessible inside scope
