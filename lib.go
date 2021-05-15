package sqlexplain

import (
	"fmt"
	"strings"

	reflect "github.com/goccy/go-reflect"
)

// Reference from : https://github.com/jmoiron/sqlx/blob/master/named.go#L313
func Explain(sql string, args ...interface{}) string {

	qb := []byte(sql)

	var sb strings.Builder
	sb.Grow(2 * len(sql))
	paramCounter := 0
	for _, v := range qb {
		if v == '?' {
			k := reflect.TypeOf(args[paramCounter])

			switch k.String() {
			case "string":
				a := args[paramCounter].(string)
				sb.WriteString("'")
				sb.WriteString(a)
				sb.WriteString("'")
			case "int":
				sb.WriteString(fmt.Sprintf("%d", args[paramCounter]))
			}

			paramCounter++
			continue
		}
		sb.WriteByte(v)
	}

	return sb.String()
}
