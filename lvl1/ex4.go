package main

import "fmt"

func main() {
	type test int

	var x test
	fmt.Println(x)
	fmt.Printf("%T\n",x)

	x = 42
	fmt.Println(x)
}
