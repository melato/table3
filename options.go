package table

import (
	"fmt"
	"io"
	"os"
)

type Options struct {
	Format string `name:"format" usage:"table format: csv|fixed"`
}

func (opt *Options) NewWriterf(w io.Writer) (Writer, error) {
	var writer Writer
	format := opt.Format
	if format == "" {
		format = "fixed"
	}
	switch format {
	case "csv":
		writer = NewCsvWriter(w)
	case "fixed":
		writer = &FixedWriter{Writer: os.Stdout}
	default:
		return nil, fmt.Errorf("unknown format: %s", opt.Format)
	}
	return writer, nil
}

func (opt *Options) NewWriter() (Writer, error) {
	return opt.NewWriterf(os.Stdout)
}
