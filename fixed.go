package table

import (
	"fmt"
	"io"
	"strings"
)

type FixedWriter struct {
	Writer    io.Writer
	NoHeaders bool
	Footer    bool
	columns   []*Column
	rows      [][]string
}

func (t *FixedWriter) Columns(columns ...*Column) {
	t.columns = columns
	if !t.NoHeaders {
		t.WriteHeader()
	}
}

func (t *FixedWriter) WriteHeader() {
	row := make([]string, len(t.columns))
	for i, col := range t.columns {
		row[i] = col.Name
	}
	t.rows = append(t.rows, row)
}

func (t *FixedWriter) WriteRow() {
	row := make([]string, len(t.columns))
	for i, col := range t.columns {
		row[i] = fmt.Sprintf("%v", col.Value())
	}
	t.rows = append(t.rows, row)
}

func (t *FixedWriter) End() {
	if t.Footer {
		t.WriteHeader()
	}
	n := len(t.columns)
	widths := make([]int, len(t.columns))
	for _, row := range t.rows {
		for j, s := range row {
			w := len(s)
			if w > widths[j] {
				widths[j] = w
			}
		}
	}
	formats := make([]string, n+1)
	for i, c := range t.columns {
		w := widths[i]
		if c.Alignment == Right {
			formats[i] = fmt.Sprintf("%%%ds", w)
		} else {
			if i == len(t.columns)-1 {
				formats[i] = "%s"
			} else {
				formats[i] = fmt.Sprintf("%%-%ds", w)
			}
		}
	}
	formats[n] = "\n"
	format := strings.Join(formats, " ")
	irow := make([]interface{}, len(t.columns))
	for _, row := range t.rows {
		for j, s := range row {
			irow[j] = s
		}
		fmt.Fprintf(t.Writer, format, irow...)
	}
}
