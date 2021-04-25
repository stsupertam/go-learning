package main

import (
	"fmt"
)

func main() {
	a := factorial(5)
	fmt.Println(a)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n*factorial(n-1)
}
