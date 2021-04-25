package main

import (
	"fmt"
)


func main() {
	a := []int{1,2,3,4,5}

	sum1 := foo(a...)
	fmt.Println(sum1)

	b := []int{10,20,30,40,50}
	sum2 := bar(b)
	fmt.Println(sum2)

}

func foo(x ...int) int {
	var sum int
	for _,val :=  range x {
		sum += val
	}
	return sum
}

func bar(x []int) int {
	var sum int
	for _,val :=  range x {
		sum += val
	}
	return sum
}