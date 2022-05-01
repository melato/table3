package table

import (
	"fmt"
	"testing"
)

type C struct {
	D string `name:"d"`
	E int    `name:"e"`
}

type S struct {
	A string `name:"a"`
	B int    `name:"b"`
	C C
}

func TestStructColumns(t *testing.T) {
	columns, err := StructColumns(&S{}, "name", func() interface{} { return nil })
	if err != nil {
		t.Fail()
	}
	if len(columns) != 4 {
		for _, col := range columns {
			fmt.Printf("%s\n", col.Name)
		}
		t.Fail()
	}
}
