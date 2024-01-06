package structs

import (
	"reflect"
)

func ConvertTo[T any](from interface{}) T {
	var to T
	fromVal := reflect.ValueOf(from)
	toVal := reflect.ValueOf(&to).Elem()

	// Check if 'fromVal' is a pointer, and if so, get the value it points to
	if fromVal.Kind() == reflect.Ptr {
		fromVal = fromVal.Elem()
	}

	if fromVal.Kind() != reflect.Struct {
		panic("ConvertTo only accepts structs or pointers to structs")
	}

	if fromVal.Type().ConvertibleTo(toVal.Type()) {
		toVal.Set(fromVal.Convert(toVal.Type()))
	} else {
		for i := 0; i < fromVal.NumField(); i++ {
			fromField := fromVal.Field(i)
			fromTypeField := fromVal.Type().Field(i)

			// Find the corresponding field in the destination struct
			if toField := toVal.FieldByName(fromTypeField.Name); toField.IsValid() && toField.CanSet() {
				if fromField.Type().AssignableTo(toField.Type()) {
					toField.Set(fromField)
				}
			}
		}
	}
	return to
}
