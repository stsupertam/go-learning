package main

import (
	"fmt"
)


func main() {
	defer foo()

	fmt.Println("Eiei eiei")

}

func foo() {
	fmt.Println("Run after na eiei")
}
