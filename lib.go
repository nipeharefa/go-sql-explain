package sqlexplain

import (
	"reflect"
)

// Reference from : https://github.com/jmoiron/sqlx/blob/master/named.go#L313
func Explain(sql string, args ...interface{}) string {

	qb := []byte(sql)

	newByte := make([]byte, 0)
	paramCounter := 0
	for _, v := range qb {
		if v == '?' {
			var s []byte
			k := reflect.TypeOf(args[paramCounter])

			switch k.String() {
			case "string":
				a := args[paramCounter].(string)
				a = "'" + a + "'"
				s = []byte(a)
			}

			newByte = append(newByte, s...)
			paramCounter++
			continue
		}
		newByte = append(newByte, v)
	}

	return string(newByte)
}
