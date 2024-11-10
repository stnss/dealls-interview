package util

import (
	"fmt"
	"strconv"
	"strings"
)

func ToString(value interface{}) string {
	switch value := value.(type) {
	case string:
		return value
	case int:
		return strconv.FormatInt(int64(value), 10)
	case int8:
		return strconv.FormatInt(int64(value), 10)
	case int16:
		return strconv.FormatInt(int64(value), 10)
	case int32:
		return strconv.FormatInt(int64(value), 10)
	case int64:
		return strconv.FormatInt(value, 10)
	case uint:
		return strconv.FormatUint(uint64(value), 10)
	case uint8:
		return strconv.FormatUint(uint64(value), 10)
	case uint16:
		return strconv.FormatUint(uint64(value), 10)
	case uint32:
		return strconv.FormatUint(uint64(value), 10)
	case uint64:
		return strconv.FormatUint(value, 10)
	case float32:
		return strconv.FormatFloat(float64(value), 'g', -1, 32)
	case float64:
		return strconv.FormatFloat(float64(value), 'g', -1, 64)
	case bool:
		return strconv.FormatBool(value)
	default:
		return fmt.Sprintf("%+v", value)
	}
}

func StringJoin(elems []string, sep, lastSep string) string {
	switch len(elems) {
	case 0:
		return ""
	case 1:
		return fmt.Sprintf("%s%s", elems[0], lastSep)
	}
	n := len(sep) * (len(elems) - 1)
	for i := 0; i < len(elems); i++ {
		n += len(elems[i])
	}

	// one, two - one and two

	var b strings.Builder
	b.Grow(n)
	b.WriteString(elems[0])
	for i := 1; i < len(elems); i++ {
		if i == len(elems)-1 && lastSep != "" {
			b.WriteString(lastSep)
			b.WriteString(elems[i])
			continue
		}

		b.WriteString(sep)
		b.WriteString(elems[i])
	}

	return b.String()
}

func SubstringAfter(src string, prefix string) string {
	// Get substring after a string.
	pos := strings.LastIndex(src, prefix)
	if pos == -1 {
		return src
	}
	adjustedPos := pos + len(prefix)
	if adjustedPos >= len(src) {
		return ""
	}
	return src[adjustedPos:]
}
