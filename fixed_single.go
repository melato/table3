package table

import (
	"fmt"
	"io"
	"os"
	"strings"
)

/*
FixedWriterSingle is uncached fixed-width writer.
It sets the column widths the same as the header widths.
*/
type FixedWriterSingle struct {
	Writer    io.Writer
	NoHeaders bool
	Footer    bool
	columns   []*Column
	widths    []int
	format    string
}

func (t *FixedWriterSingle) Columns(columns ...*Column) {
	t.columns = columns
	if !t.NoHeaders {
		t.WriteHeader()
	}
}

func (t *FixedWriterSingle) WriteHeader() {
	n := len(t.columns)
	widths := make([]int, n)
	formats := make([]string, n+1)
	formats[n] = "\n"

	row := make([]string, len(t.columns))
	for i, col := range t.columns {
		row[i] = col.Name
		w := width(col.Name)
		widths[i] = w
		var format string
		if col.Alignment == Right {
			format = fmt.Sprintf("%%%ds", w)
		} else {
			if i == len(t.columns)-1 {
				format = "%s"
			} else {
				format = fmt.Sprintf("%%-%ds", w)
			}
		}
		formats[i] = format
	}
	t.widths = widths
	t.format = strings.Join(formats, " ")
	t.write(row)
}

func (t *FixedWriterSingle) WriteRow() {
	row := make([]string, len(t.columns))
	for i, col := range t.columns {
		row[i] = fmt.Sprintf("%v", col.Value())
	}
	t.write(row)
}

func (t *FixedWriterSingle) End() {
}

func (t *FixedWriterSingle) write(row []string) {
	writer := t.Writer
	if writer == nil {
		writer = os.Stdout
	}
	irow := make([]interface{}, len(t.columns))
	for j, s := range row {
		irow[j] = s
	}
	fmt.Fprintf(writer, t.format, irow...)
}
