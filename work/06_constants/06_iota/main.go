package main

import "fmt"

const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1*10)
	MB = 1 << (iota * 10) // 1 << (2*10)
)

func main() {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t%d\n", KB, KB)
	fmt.Printf("%b\t%d\n", MB, MB)
}
