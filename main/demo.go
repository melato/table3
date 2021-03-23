package main

import (
	"fmt"
	"os"

	"melato.org/table3"
)

func main() {
	writer := &table.FixedWriter{Writer: os.Stdout}
	var a string
	var b float32
	writer.Columns(
		table.NewColumn("a", func() interface{} { return a }),
		table.NewColumn("b", func() interface{} { return fmt.Sprintf("%10.2f", b) }).Align(table.Right),
	)
	for i := 0; i < 50; i += 5 {
		a = fmt.Sprintf("a%d", i)
		b = float32(i*1000) + 0.01*float32(i)
		writer.WriteRow()
	}
	writer.End()
}
