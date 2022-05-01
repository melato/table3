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

func (opt *FullOptions) PrintSlice(slice interface{}, nameTag string) error {
	sliceType := reflect.TypeOf(slice)
	if sliceType.Kind() != reflect.Slice {
		return fmt.Errorf("not a slice: %v", sliceType)
	}
	sliceValue := reflect.ValueOf(slice)
	proto := reflect.Zero(sliceType.Elem())
	var v interface{}
	columns, err := StructColumns(proto.Interface(), nameTag, func() interface{} { return v })
	w, err := opt.NewWriter(columns...)
	if err != nil {
		return err
	}
	n := sliceValue.Len()
	for i := 0; i < n; i++ {
		v = sliceValue.Index(i).Interface()
		w.WriteRow()
	}
	w.End()
	return nil
}
