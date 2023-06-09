package v2

import (
	"reflect"
)

func walk(x interface{}, fn func(input string)) {
	val := reflect.ValueOf(x)
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.String {
			fn(field.String())
		} else if field.Kind() == reflect.Struct {
			walk(field.Interface(), fn)
		}
	}
}
