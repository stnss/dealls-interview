package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestToString(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{"string", "hello", "hello"},
		{"int", 42, strconv.FormatInt(42, 10)},
		{"int8", int8(127), strconv.FormatInt(127, 10)},
		{"int16", int16(32767), strconv.FormatInt(32767, 10)},
		{"int32", int32(2147483647), strconv.FormatInt(2147483647, 10)},
		{"int64", int64(9223372036854775807), strconv.FormatInt(9223372036854775807, 10)},
		{"uint", uint(42), strconv.FormatUint(42, 10)},
		{"uint8", uint8(255), strconv.FormatUint(255, 10)},
		{"uint16", uint16(65535), strconv.FormatUint(65535, 10)},
		{"uint32", uint32(4294967295), strconv.FormatUint(4294967295, 10)},
		{"uint64", uint64(18446744073709551615), strconv.FormatUint(18446744073709551615, 10)},
		{"float32", float32(3.14), strconv.FormatFloat(3.14, 'g', -1, 32)},
		{"float64", float64(2.718281828459), strconv.FormatFloat(2.718281828459, 'g', -1, 64)},
		{"bool", true, strconv.FormatBool(true)},
		{"slice", []int{1, 2, 3}, fmt.Sprintf("%+v", []int{1, 2, 3})},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, ToString(tt.input))
		})
	}
}

func TestStringJoin(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		sep      string
		lastSep  string
		expected string
	}{
		{"empty slice", []string{}, ", ", " and ", ""},
		{"single element", []string{"one"}, ", ", " and ", "one and "},
		{"two elements", []string{"one", "two"}, ", ", " and ", "one and two"},
		{"multiple elements", []string{"one", "two", "three", "four"}, ", ", " and ", "one, two, three and four"},
		{"no last separator", []string{"one", "two", "three", "four"}, ", ", "", "one, two, three, four"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, StringJoin(tt.input, tt.sep, tt.lastSep))
		})
	}
}

func TestSubstringAfter(t *testing.T) {
	tests := []struct {
		name     string
		src      string
		prefix   string
		expected string
	}{
		{"prefix found", "hello world", "hello ", "world"},
		{"prefix not found", "hello world", "foo", "hello world"},
		{"prefix at the end", "hello world", "world", ""},
		{"prefix twice", "bar foo foo", "foo ", "foo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, SubstringAfter(tt.src, tt.prefix))
		})
	}
}
