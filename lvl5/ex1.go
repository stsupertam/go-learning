package main

import (
	"fmt"
)

type person struct{
	firstname string
	lastname string
	iceCream[] string
}
func main() {
	a := person{
		firstname: "Test1",
		lastname: "Test2",
		iceCream: []string{"Chocolate","Vanilla"},
	}
	fmt.Println(a)
	for _,v := range(a.iceCream) {
		fmt.Println(v)
	}
	b := person{
		firstname: "A",
		lastname: "B",
		iceCream: []string{"Lemon","Coffee"},
	}
	fmt.Println(b)
	for _,v := range(b.iceCream) {
		fmt.Println(v)
	}


}
