package main

import (
	"fmt"
)

const (
	x = 42
	y int = 43
)
func main() {
	fmt.Printf("%T\n",x)
	fmt.Printf("%T\n",y)
	fmt.Println(x,y)
}
