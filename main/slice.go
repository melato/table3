package main

import (
	"fmt"

	table "melato.org/table3"
)

type T struct {
	A string `name:"a"`
	B int    `name:"b"`
}

func main() {
	list := []*T{
		&T{"a", 1},
		&T{"b", 2},
	}
	var options table.FullOptions
	err := options.PrintSlice(list, "name")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
