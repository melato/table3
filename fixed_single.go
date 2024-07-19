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
	rows      [][]string
	count     int
}

func (t *FixedWriterSingle) Columns(columns ...*Column) {
	t.columns = columns
	if !t.NoHeaders {
		t.WriteHeader()
	}
}

type fixedFormat struct {
}

func (t *fixedFormat) computeWidths(rows [][]string) []int {
	if len(rows) == 0 {
		return nil
	}
	n := len(rows[0])
	widths := make([]int, n)
	for _, row := range rows {
		for j, s := range row {
			w := width(s)
			if w > widths[j] {
				widths[j] = w
			}
		}
	}
	return widths
}

func (t *FixedWriterSingle) computeFormats() {
	var format fixedFormat
	widths := format.computeWidths(t.rows)

	n := len(widths)
	formats := make([]string, n+1)
	formats[n] = "\n"

	for i, w := range widths {
		var format string
		if i < len(t.columns) && t.columns[i].Alignment == Right {
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
}

func (t *FixedWriterSingle) addRow(row []string) {
	t.count += 1
	switch t.count {
	case 1:
		t.rows = append(t.rows, row)
	case 2:
		t.rows = append(t.rows, row)
		t.computeFormats()
		for _, row := range t.rows {
			t.write(row)
		}
	default:
		t.write(row)
	}
}

func (t *FixedWriterSingle) WriteHeader() {
	row := make([]string, len(t.columns))
	for i, col := range t.columns {
		row[i] = col.Name
	}
	t.addRow(row)
}

func (t *FixedWriterSingle) WriteRow() {
	row := make([]string, len(t.columns))
	for i, col := range t.columns {
		row[i] = fmt.Sprintf("%v", col.Value())
	}
	t.addRow(row)
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
