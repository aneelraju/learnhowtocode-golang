package main

import "fmt"

func main() {
	x := 13 % 3
	fmt.Println(x)

	for i := 0; i < 100; i++ {
		if i%2 == 1 {
			fmt.Printf("%d is Odd\n", i)
		} else {
			fmt.Printf("%d is Even\n", i)
		}
	}
}
