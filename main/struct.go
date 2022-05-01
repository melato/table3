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

func printList[T comparable](list []*T) error {
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
	var err error
	if true {
		list1 := []*S{
			&S{"a", 1},
			&S{"b", 2},
		}
		err = printList[S](list1)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
	if true {
		list2 := []*C{
			&C{S{"x", 3}, "a", 1},
			&C{S{"y", 4},"b", 2},
		}
		err = printList[C](list2)
		if err != nil {
			fmt.Printf("%v\n", err)
		}
	}
}
