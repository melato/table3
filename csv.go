package table

import (
	"encoding/csv"
	"fmt"
	"io"
)

type CsvWriter struct {
	Writer  *csv.Writer
	columns []*Column
	row     []string
}

func NewCsvWriter(w io.Writer) Writer {
	return &CsvWriter{Writer: csv.NewWriter(w)}
}

func (t *CsvWriter) End() {
	t.Writer.Flush()
}

func (t *CsvWriter) Columns(columns ...*Column) {
	t.columns = columns
	t.WriteHeader()
}

func (t *CsvWriter) WriteHeader() {
	t.row = make([]string, len(t.columns))
	for i, col := range t.columns {
		t.row[i] = col.Name
	}
	t.Writer.Write(t.row)
}

func (t *CsvWriter) WriteRow() {
	for i, col := range t.columns {
		t.row[i] = fmt.Sprintf("%v", col.Value())
	}
	t.Writer.Write(t.row)
}
