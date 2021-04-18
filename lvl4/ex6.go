package main

import (
	"fmt"
)

func main() {
	x := []string{"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming" }
	fmt.Println(len(x))
	fmt.Println(cap(x))

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}

	y := make([]string, 50, 500)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// put a value into each of the 50 positions in the length of the slice
	// we are putting values into the 50 positions that are the length of the slice
	for i := 0; i < 50; i++ {
		y[i] = fmt.Sprintf("Position %d", i)
	}

	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	// append values to the slice grows the length of the slice
	// the underlying array "capacity" of 500 is the same
	y = append(y, ` Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(cap(y))

	for i := 0; i < len(y); i++ {
		fmt.Println(i, y[i])
	}
}
