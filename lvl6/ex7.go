package main

import (
	"fmt"
)

func main() {
	a := func() {
		fmt.Print("Annonymous")
	}
	a()
}

