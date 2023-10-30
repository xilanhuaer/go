package orm

import (
	"errors"
	"reflect"
)

// model is a ptr
func Create(model interface{}, name string) error {
	dest := reflect.TypeOf(model)
	value := reflect.ValueOf(model)
	if dest.Kind() != reflect.Ptr || value.Kind() != reflect.Ptr {
		return errors.New("model must be a pointer to a struct")
	}
	dest = dest.Elem()
	value = value.Elem()
	for i := 0; i < dest.NumField(); i++ {
		fieldName := dest.Field(i).Name
		fieldType := dest.Field(i).Type
		if fieldType.Kind() != reflect.String {
			continue
		}
		switch fieldName {
		case "Creator", "Updator":
			value.FieldByName(fieldName).SetString(name)
		}
	}
	return nil
}
