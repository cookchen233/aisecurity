package structs

import (
	"encoding/json"
	"entgo.io/ent"
	"fmt"
	"hash/crc32"
	"reflect"
)

func ConvertTo[T any](from interface{}) T {
	var to T
	fromVal := reflect.ValueOf(from)
	toType := reflect.TypeOf(to)

	// Create a new instance of the type 'to' points to if 'to' is a pointer
	var toVal reflect.Value
	if toType.Kind() == reflect.Ptr {
		toVal = reflect.New(toType.Elem()).Elem()
	} else {
		toVal = reflect.New(toType).Elem()
	}

	// Ensure fromVal is not a pointer
	if fromVal.Kind() == reflect.Ptr && !fromVal.IsNil() {
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

	// Return the converted value, converting back to pointer if necessary
	if toType.Kind() == reflect.Ptr {
		return toVal.Addr().Interface().(T)
	}
	return toVal.Interface().(T)
}

// IFilter represents qeury parameters
type IFilter interface {
	GetPage() int
	GetLimit() int
	GetOffset() int
	GetID() int
}

type StandardFilter struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
	ID    int `form:"id"`
}

func (f *StandardFilter) GetPage() int { return f.Page }

func (f *StandardFilter) GetLimit() int { return f.Limit }

func (f *StandardFilter) GetOffset() int {
	page := min(1000, max(1, f.Page))
	limit := min(10000, max(1, f.Limit))
	return (page - 1) * limit
}

func (f *StandardFilter) GetID() int { return f.ID }

type IEntity interface {
	Value(name string) (ent.Value, error)
}

type StandardIEntity struct{}

func (ent *StandardIEntity) Value(name string) (ent.Value, error) {
	return nil, nil
}

// IPost represents body parameters
type IPost interface {
	GetID() int
}

type StandardPost struct {
	ID int `json:"id"`
}

func (p *StandardPost) GetID() int {
	return p.ID
}

// IResponse represents output data
type IResponse interface {
	Hash() string
}

type StandardResponse struct {
	ID int `json:"id"`
}

func (p *StandardResponse) Hash() string {
	bytes, err := json.Marshal(p)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%08x", crc32.ChecksumIEEE(bytes))
}
