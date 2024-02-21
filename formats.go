package table

import (
	"fmt"
	"io"
)

var formats = make(map[string]func(io.Writer) Writer)

func SetFormat(name string, fn func(io.Writer) Writer) {
	formats[name] = fn
}

func NewWriterf(format string, w io.Writer) (Writer, error) {
	f, ok := formats[format]
	if !ok {
		return nil, fmt.Errorf("unknown table format: %s", format)
	}
	return f(w), nil
}

func init() {
	SetFormat("csv", NewCsvWriter)
	SetFormat("fixed", func(w io.Writer) Writer { return &FixedWriter{Writer: w} })
	SetFormat("fixed-onepass", func(w io.Writer) Writer { return &FixedWriterSingle{Writer: w} })
	SetFormat("json", func(w io.Writer) Writer { return &JsonWriter{Writer: w} })
}
