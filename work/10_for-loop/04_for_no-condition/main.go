package main

import "fmt"

func main() {
	i := 0
	// runs until interrupted in command prompt (ctrl+c)
	for {
		fmt.Println(i)
		i++
	}
}
