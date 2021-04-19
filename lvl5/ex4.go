package main

import (
	"fmt"
)

func main() {
	var a struct {
		test1 string
		test2 string
	}

	a.test1 = "Gundam"
	a.test2 = "Amuro Ray"

	fmt.Println(a)

	s := struct {
		a string
		b string
	}{
		a: "test1",
		b: "test2",
	}

	fmt.Println(s)
}
