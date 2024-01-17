package structs

import (
	"aisecurity/utils"
	"encoding/json"
	"entgo.io/ent"
	"fmt"
	"github.com/gogf/gf/v2/util/gconv"
	"go.uber.org/zap"
	"hash/crc32"
	"reflect"
)

func ConvertTo2[T any](from interface{}) T {
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

func ConvertListTo[Tdata any, Tto any](list []Tdata) []*Tto {
	list2 := make([]*Tto, len(list))
	for i, v := range list {
		list2[i] = ConvertTo[Tdata, Tto](v)
	}
	return list2
}

func ConvertTo[Tdata any, Tto any](data Tdata) *Tto {
	data2 := new(Tto)
	err := gconv.Struct(data, data2)
	if err != nil {
		utils.Logger.Error("failed converting", zap.Error(err))
	}
	return data2
}

// IFilter represents qeury parameters
type IFilter interface {
	GetPage() int
	SetPage(page int)
	GetLimit() int
	SetLimit(limit int)
	GetOffset() int
	GetID() int
}

type StandardFilter struct {
	// for generic filter
	Page  int `form:"page"`
	Limit int `form:"limit"`
	ID    int `form:"id"`

	// for specific filter, some filters may not need.
	Name  string `form:"name"`
	Title string `form:"title"`
}

func (f *StandardFilter) GetPage() int     { return min(1000, max(1, f.Page)) }
func (f *StandardFilter) SetPage(page int) { f.Page = page }

func (f *StandardFilter) GetLimit() int      { return min(10000, max(1, f.Limit)) }
func (f *StandardFilter) SetLimit(limit int) { f.Limit = limit }

func (f *StandardFilter) GetOffset() int { return (f.GetPage() - 1) * f.GetLimit() }

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
