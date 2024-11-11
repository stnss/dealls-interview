package util

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

func IsSameType(src, dest interface{}) bool {
	return reflect.TypeOf(src) == reflect.TypeOf(dest)
}

// StructToMap converts a struct to a map using the struct's tags.
// StructToMap uses tags on struct fields to decide which fields to add to the
// returned map.
func StructToMap(src interface{}, tag string) (map[string]interface{}, error) {
	out := map[string]interface{}{}
	v := reflect.ValueOf(src)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// we only accept structs
	if v.Kind() != reflect.Struct {
		return out, fmt.Errorf("only accepted %s, got %s", reflect.Struct.String(), v.Kind().String())
	}

	typ := v.Type()
	for i := 0; i < v.NumField(); i++ {
		// gets us a StructField
		fi := typ.Field(i)

		//field := reflectValue.Field(i).Interface()
		if !fi.IsExported() {
			continue
		}

		tagsv := strings.Split(fi.Tag.Get(tag), ",")
		prefix := strings.Split(tagsv[len(tagsv)-1], "=")
		if tagsv[0] != "" && fi.PkgPath == "" {

			if isNil(v.Field(i).Interface()) {
				continue
			}

			// skip if omitempty
			if (len(tagsv) > 1 && tagsv[1] == "omitempty") && IsEmptyValue(v.Field(i).Interface()) {
				continue
			}

			if isTime(v.Field(i)) {
				if timeIsZero(v.Field(i)) && (len(tagsv) > 1 && tagsv[1] == "omitempty") {
					continue
				}
			}

			if v.Field(i).Kind() == reflect.Struct {
				continue
			}

			col := tagsv[0]

			if len(prefix) > 1 && prefix[0] == "prefix" && prefix[1] != "" {
				col = prefix[1] + col
			}

			if InArray("ne", tagsv) {
				col = fmt.Sprintf("%s !", col)
			}
			// set key value of map interface output
			out[col] = v.Field(i).Interface()
		}

		if tagsv[0] == "" && v.Field(i).Kind() == reflect.Struct {
			x, err := StructToMap(v.Field(i).Interface(), tag)
			if err != nil {
				return out, err
			}

			for y, z := range x {
				out[y] = z
			}
		}
	}

	return out, nil
}

func isNil(i interface{}) bool {
	if i == nil || (reflect.ValueOf(i).Kind() == reflect.Ptr && reflect.ValueOf(i).IsNil()) {
		return true
	}

	return false
}

func isTime(obj reflect.Value) bool {
	_, ok := obj.Interface().(time.Time)
	if ok {
		return ok
	}

	_, ok = obj.Interface().(*time.Time)

	return ok
}

func timeIsZero(obj reflect.Value) bool {
	t, ok := obj.Interface().(time.Time)
	if ok {
		return t.IsZero()
	}

	t2, ok := obj.Interface().(*time.Time)
	if ok {
		return false
	}

	return t2 == nil
}

func CopyStruct(src, dst any, tag string) error {
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)

	// If src or dst is a pointer, get the underlying value.
	if srcVal.Kind() == reflect.Ptr {
		srcVal = srcVal.Elem()
	}
	if dstVal.Kind() == reflect.Ptr {
		dstVal = dstVal.Elem()
	}

	// Ensure both src and dst are structs.
	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return fmt.Errorf("src and dst must be structs or pointers to structs")
	}

	// Map to store dst field by JSON tag for quick lookup
	dstFieldMap := make(map[string]reflect.Value)

	// Populate dstFieldMap with dst fields that have json tags.
	for i := 0; i < dstVal.NumField(); i++ {
		dstField := dstVal.Type().Field(i)
		if jsonTag, ok := dstField.Tag.Lookup(tag); ok {
			tagName := strings.Split(jsonTag, ",")[0] // Get the tag before any options like omitempty
			if tagName != "-" {
				dstFieldMap[tagName] = dstVal.Field(i)
			}
		}
	}

	// Iterate over src fields and copy to dst where tags match.
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Type().Field(i)
		srcFieldValue := srcVal.Field(i)

		// Skip unexported fields (fields that start with lowercase letters)
		if !srcFieldValue.CanInterface() {
			continue
		}

		// Check if the field has a JSON tag and is not "-"
		if jsonTag, ok := srcField.Tag.Lookup(tag); ok {
			tagName := strings.Split(jsonTag, ",")[0]
			omitempty := strings.Contains(jsonTag, "omitempty")

			// Check for zero value and omitempty tag
			if omitempty && isZeroValue(srcFieldValue) {
				continue // Skip this field if it's zero and has omitempty tag
			}

			// If there's a corresponding field in dst, set its value.
			if dstField, exists := dstFieldMap[tagName]; exists && dstField.CanSet() {
				if srcFieldValue.Kind() == reflect.Ptr && !srcFieldValue.IsNil() {
					// Dereference the pointer
					srcFieldValue = srcFieldValue.Elem()
				}

				// Handle pointer types by ensuring compatibility between source and destination
				if srcFieldValue.Type().AssignableTo(dstField.Type()) {
					dstField.Set(srcFieldValue)
				} else if srcFieldValue.Kind() == reflect.Ptr && dstField.Kind() == reflect.Ptr {
					if !srcFieldValue.IsNil() && dstField.CanSet() {
						dstField.Set(srcFieldValue)
					}
				} else if dstField.Kind() == reflect.Ptr && dstField.IsNil() {
					if srcFieldValue.CanAddr() {
						dstField.Set(srcFieldValue.Addr())
					}
				}
			}
		}
	}
	return nil
}

// isZeroValue checks if a value is the zero value for its type.
func isZeroValue(v reflect.Value) bool {
	if !v.IsValid() {
		return true
	}

	if v.Kind() == reflect.Ptr || v.Kind() == reflect.Interface {
		return v.IsNil()
	}

	switch v.Kind() {
	case reflect.String, reflect.Slice, reflect.Map, reflect.Array, reflect.Chan:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Struct:
		return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
	default:
		return false
	}
}
