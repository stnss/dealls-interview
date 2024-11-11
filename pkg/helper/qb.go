package helper

import (
	"fmt"
	"github.com/stnss/dealls-interview/pkg/util"
	"strings"
)

func StructQueryInsert(param interface{}, tableName, tag string, returningID bool) (string, []interface{}, error) {
	var (
		keys   []string
		values []interface{}
		numArr []string
	)

	resMap, err := util.StructToMap(param, tag)
	if err != nil {
		return "", nil, err
	}

	for k, v := range resMap {
		keys = append(keys, k)
		values = append(values, v)
		numArr = append(numArr, "?")
	}

	q := ""
	if returningID {
		q = fmt.Sprintf(`
		INSERT INTO
			%s
		(
			%s
		)
		VALUES
		(
			%s
		)
		RETURNING id
	`, tableName, strings.Join(keys, ","), strings.Join(numArr, ","))
	} else {
		q = fmt.Sprintf(`
		INSERT INTO
			%s
		(
			%s
		)
		VALUES
		(
			%s
		)
	`, tableName, strings.Join(keys, ","), strings.Join(numArr, ","))
	}

	return q, values, nil
}
