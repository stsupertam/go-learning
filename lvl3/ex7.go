package main

import (
	"fmt"
)

func main() {
	if i := 5; i == 6 {
		fmt.Println("I == 6")
	} else if i == 5 {
		fmt.Println("I == 5")
	}
}
