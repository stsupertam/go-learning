package main

import (
	"fmt"
)

func main() {
    var m = map[string][]string {
        "bond_james": []string{`Shaken, not stirred`, `Martinis`, `Women`},
        "moneypenny_miss": []string{`James Bond`, `Literature`, `Computer Science`},
        "no_dr": []string{`Being evil`, `Ice cream`, `Sunsets`},
    }

	m["test1_test2"] = []string{`Test1`, `Test2`, `Test3`}
	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, v2 := range v {
			fmt.Println("\t", i, v2)
		}
	}
}
