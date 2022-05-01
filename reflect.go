package table

import (
	"fmt"
	"reflect"
)

func StructColumns(proto interface{}, nameTag string, row func() interface{}) ([]*Column, error) {
	vType := reflect.TypeOf(proto)
	var isPointer bool
	if vType.Kind() == reflect.Pointer {
		isPointer = true
		vType = vType.Elem()
	}
	if vType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not a struct: %v", vType)
	}
	var columns []*Column
	n := vType.NumField()
	for i := 0; i < n; i++ {
		f := vType.Field(i)
		if !f.IsExported() {
			continue
		}
		var name string
		if nameTag != "" {
			name = f.Tag.Get(nameTag)
			if name == "-" {
				continue
			}
		}
		if name == "" {
			name = f.Name
		}
		j := i
		columns = append(columns, NewColumn(name, func() interface{} {
			v := row()
			vValue := reflect.ValueOf(v)
			if isPointer {
				vValue = vValue.Elem()
			}
			fv := vValue.Field(j)
			return fv.Interface()
		}))
	}
	return columns, nil
}
