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
	b := person{
		firstname: "A",
		lastname: "B",
		iceCream: []string{"Lemon","Coffee"},
	}

	persons := map[string]person{}
	persons[a.lastname] = a
	persons[b.lastname] = b

	for _, v := range persons {
		fmt.Println(v)
		for _, favor := range v.iceCream {
			fmt.Println(favor)
		} 
	}

}
