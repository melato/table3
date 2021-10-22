package table

import (
	"fmt"
	"os"
)

type Options struct {
	Format string `name:"format" usage:"table format: csv|fixed"`
}

func (opt *Options) NewWriter() (Writer, error) {
	var w Writer
	format := opt.Format
	if format == "" {
		format = "fixed"
	}
	switch format {
	case "csv":
		w = NewCsvWriter(os.Stdout)
	case "fixed":
		w = &FixedWriter{Writer: os.Stdout}
	default:
		return nil, fmt.Errorf("unknown format: %s", opt.Format)
	}
	return w, nil
}
