package main

import (
	"fmt"

	table "melato.org/table3"
)

type S struct {
	A string `name:"a"`
	B int    `name:"b"`
}

type C struct {
	S
	D string `name:"d"`
	E int    `name:"e"`
}

func main() {
	list := []*C{
		&C{S{"x", 3}, "a", 1},
		&C{S{"y", 4}, "b", 2},
	}
	var options table.FullOptions
	err := options.PrintSlice(list, "name")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
