package table

import (
	"fmt"
	"reflect"
)

func structColumns(vType reflect.Type, nameTag string, row func() interface{}) ([]*Column, error) {
	var columns []*Column
	n := vType.NumField()
	for i := 0; i < n; i++ {
		f := vType.Field(i)
		if !f.IsExported() {
			continue
		}
		switch f.Type.Kind() {
		case reflect.Chan, reflect.Func, reflect.Pointer:
			continue
		case reflect.Struct:
			j := i
			fcols, err := structColumns(f.Type, nameTag, func() interface{} {
				v := row()
				vValue := reflect.ValueOf(v)
				if vValue.Type().Kind() == reflect.Pointer {
					vValue = vValue.Elem()
				}
				fv := vValue.Field(j)
				return fv.Interface()
			})
			if err != nil {
				return nil, err
			}
			columns = append(columns, fcols...)
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
			if vValue.Type().Kind() == reflect.Pointer {
				vValue = vValue.Elem()
			}
			fv := vValue.Field(j)
			return fv.Interface()
		}))
	}
	return columns, nil
}

func StructColumns(proto interface{}, nameTag string, row func() interface{}) ([]*Column, error) {
	vType := reflect.TypeOf(proto)
	//var isPointer bool
	if vType.Kind() == reflect.Pointer {
		//isPointer = true
		vType = vType.Elem()
	}
	if vType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("not a struct: %v", vType)
	}
	return structColumns(vType, nameTag, row)
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
