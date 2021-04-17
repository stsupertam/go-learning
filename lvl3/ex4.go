package main

import (
	"fmt"
)

func main() {
	var i int = 1995
	for {
		if i > 2021 {
			break
		}
		fmt.Println(i)
		i++
	}
}
