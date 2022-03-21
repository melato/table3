package table

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Options struct {
	Format string `name:"format" usage:"table format: csv|fixed"`
}

type FullOptions struct {
	Options
	Columns string `name:"columns" usage:"comma-separated list of columns to use"`
	columns []string
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

func (opt *Options) NewWriterWithColumns(columnSpec string, allColumns ...*Column) (Writer, error) {
	w, err := opt.NewWriterf(os.Stdout)
	if err != nil {
		return nil, err
	}
	columns := allColumns
	if columnSpec != "" {
		columns, err = SelectColumns(allColumns, strings.Split(columnSpec, ",")...)
		if err != nil {
			PrintColumns(allColumns)
			return nil, err
		}
	}
	w.Columns(columns...)
	return w, nil
}

func (opt *FullOptions) NewWriter(allColumns ...*Column) (Writer, error) {
	opt.columns = strings.Split(opt.Columns, ",")
	w, err := opt.NewWriterf(os.Stdout)
	if err != nil {
		return nil, err
	}
	columns := allColumns
	if opt.Columns != "" {
		columns, err = SelectColumns(allColumns, opt.columns...)
		if err != nil {
			PrintColumns(allColumns)
			return nil, err
		}
	}
	w.Columns(columns...)
	return w, nil
}

func (opt *FullOptions) HasColumn(name string) bool {
	for _, col := range opt.columns {
		if col == name {
			return true
		}
	}
	return false
}
