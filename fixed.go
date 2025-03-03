package table

import (
	"fmt"
	"io"
	"os"
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

// width - return the number of unicode characters in the string.
func width(s string) int {
	var w int
	for _, _ = range s {
		w++
	}
	return w
}

func (t *FixedWriter) End() {
	if t.Writer == nil {
		t.Writer = os.Stdout
	}
	if t.Footer {
		t.WriteHeader()
	}
	n := len(t.columns)
	widths := make([]int, len(t.columns))
	for _, row := range t.rows {
		for j, s := range row {
			w := width(s)
			if w > widths[j] {
				widths[j] = w
			}
		}
	}
	formats := make([]string, n)
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
	format := strings.Join(formats, " ") + "\n"
	irow := make([]interface{}, len(t.columns))
	for _, row := range t.rows {
		for j, s := range row {
			irow[j] = s
		}
		fmt.Fprintf(t.Writer, format, irow...)
	}
}
