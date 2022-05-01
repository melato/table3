package main

import (
	"fmt"

	table "melato.org/table3"
)

type T struct {
	A string `name:"a"`
	B int    `name:"b"`
}

func printList(list []*T) error {
	var options table.FullOptions
	var v *T
	columns, err := table.StructColumns(v, "name", func() interface{} { return v })
	if err != nil {
		return err
	}
	w, err := options.NewWriter(columns...)
	if err != nil {
		return err
	}
	for _, v = range list {
		w.WriteRow()
	}
	w.End()
	return nil
}

func main() {
	list := []*T{
		&T{"a", 1},
		&T{"b", 2},
	}
	err := printList(list)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
}
