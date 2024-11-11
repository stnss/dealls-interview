package util

import (
	"github.com/gofiber/fiber/v2/log"
	"reflect"
)

// InArray check if an element is exist in the array
func InArray(val interface{}, array interface{}) bool {
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				return true
			}
		}
	default:
		log.Error("Unhandled array type.")
	}
	return false
}
