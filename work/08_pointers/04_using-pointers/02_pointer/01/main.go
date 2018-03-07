package main

import "fmt"

func zero(z *int) {
	*z = 0
}

func main() {
	x := 5
	zero(&x)       // pass by value
	fmt.Println(x) // x is 0
}
