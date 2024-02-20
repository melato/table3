package table

import (
	"encoding/json"
	"io"
)

type JsonWriter struct {
	Writer  io.Writer
	columns []*Column
	rows    []map[string]any
}

func (t *JsonWriter) End() {
	data, _ := json.MarshalIndent(t.rows, "", " ")
	t.Writer.Write(data)
}

func (t *JsonWriter) Columns(columns ...*Column) {
	t.columns = columns
}

func (t *JsonWriter) WriteRow() {
	row := make(map[string]any)
	for _, col := range t.columns {
		row[col.Name] = col.Value()
	}
	t.rows = append(t.rows, row)
}
