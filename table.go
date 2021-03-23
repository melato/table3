// Package table defines an interface for defining and using table columns for printing,
// and two implementations:  fixed-width and CSV.
package table

type Alignment int

const (
	Left  Alignment = 0
	Right Alignment = 1
)

// Specifies a column, including static column description, and current data.
type Column struct {
	// The display name of the column
	Name string
	// Internal column identifier, defaults to Name
	Identifier string
	// The value of the column for the current row, when called by Writer.Write().  It will be converted to a string using fmt %v.
	Value func() interface{}
	// Alignment: Left or Right
	Alignment Alignment
	// optional column description
	Description string
}

func NewColumn(name string, value func() interface{}) *Column {
	return (&Column{Name: name, Value: value}).Id(name)
}

func (c *Column) Align(alignment Alignment) *Column {
	c.Alignment = alignment
	return c
}

func (c *Column) Desc(description string) *Column {
	c.Description = description
	return c
}

func (c *Column) Id(id string) *Column {
	c.Identifier = id
	return c
}

// Interface for writing tabular data
type Writer interface {
	// Specify the columns
	Columns(...*Column)
	// Write the header.  May be called multiple times.
	//WriteHeader()
	// Write the current row.
	WriteRow()
	// End writing.  Flush or close files.
	End()
}
