package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSameType(t *testing.T) {
	tests := []struct {
		name     string
		src      interface{}
		dest     interface{}
		expected bool
	}{
		{"both nil", nil, nil, true},
		{"one nil, one int", nil, 1, false},
		{"int and int", 1, 2, true},
		{"int and int64", 1, int64(2), false},
		{"float32 and float64", float32(1.1), float64(2.2), false},
		{"float64 and float64", float64(1.1), float64(2.2), true},
		{"string and string", "hello", "world", true},
		{"string and int", "hello", 1, false},
		{"array and array of same type", [2]int{1, 2}, [2]int{3, 4}, true},
		{"array and array of different size", [2]int{1, 2}, [3]int{1, 2, 3}, false},
		{"slice and slice of same type", []int{1, 2}, []int{3, 4}, true},
		{"slice and slice of different type", []int{1, 2}, []string{"a", "b"}, false},
		{"map and map of same type", map[string]int{"one": 1}, map[string]int{"two": 2}, true},
		{"map and map of different type", map[string]int{"one": 1}, map[int]string{1: "one"}, false},
		{"struct and struct of same type", struct{ Name string }{"Alice"}, struct{ Name string }{"Bob"}, true},
		{"struct and struct of different field type", struct{ Name string }{"Alice"}, struct{ Name int }{42}, false},
		{"pointer and pointer of same type", &[]int{1, 2}, &[]int{3, 4}, true},
		{"pointer and pointer of different type", &[]int{1, 2}, &[]string{"a", "b"}, false},
		{"func and func of same signature", func() {}, func() {}, true},
		{"func and func of different signature", func() {}, func(int) {}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, IsSameType(tt.src, tt.dest))
		})
	}
}
