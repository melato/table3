package table

import (
	"io"
	"os"
	"strings"
)

type Options struct {
	Format string `name:"format" usage:"table format: csv|fixed|json"`
}

type FullOptions struct {
	Options
	Columns string `name:"columns" usage:"comma-separated list of columns to use"`
	columns []string
}

func (opt *Options) NewWriterf(w io.Writer) (Writer, error) {
	format := opt.Format
	if format == "" {
		format = "fixed"
	}
	return NewWriterf(format, w)
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

func (opt *FullOptions) NewWriterf(out io.Writer, allColumns ...*Column) (Writer, error) {
	opt.columns = strings.Split(opt.Columns, ",")
	w, err := opt.Options.NewWriterf(out)
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

func (opt *FullOptions) NewWriter(allColumns ...*Column) (Writer, error) {
	return opt.NewWriterf(os.Stdout, allColumns...)
}

func (opt *FullOptions) HasColumn(name string) bool {
	for _, col := range opt.columns {
		if col == name {
			return true
		}
	}
	return false
}
