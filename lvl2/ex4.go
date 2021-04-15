package main

import (
	"fmt"
)


func main() {
	var x int = 42
	fmt.Printf("%d\t%b\t%x\n", x, x, x)
	var y = x << 1
	fmt.Printf("%d\t%b\t%x\n", y, y, y)
}
