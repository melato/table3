package table

import (
	"errors"
	"fmt"
	"strings"
)

func PrintColumns(columns []*Column) {
	var idLen int
	for _, c := range columns {
		w := len(c.Identifier)
		if w > idLen {
			idLen = w
		}
	}
	fmt.Printf("Available columns:\n")
	for _, c := range columns {
		fmt.Printf("  %s:%*s %s\n", c.Identifier, idLen-len(c.Identifier), "", c.Description)
	}
}

func SelectColumns(allColumns []*Column, ids ...string) ([]*Column, error) {
	cm := make(map[string]*Column)
	for _, c := range allColumns {
		cm[c.Identifier] = c
	}

	var cols []*Column
	var unknown []string
	for _, id := range ids {
		col, found := cm[id]
		if !found {
			unknown = append(unknown, id)
			continue
		}
		if len(unknown) == 0 {
			cols = append(cols, col)
		}
	}
	if len(unknown) > 0 {
		return nil, errors.New("unknown columns: " + strings.Join(unknown, ","))
	}
	return cols, nil
}
