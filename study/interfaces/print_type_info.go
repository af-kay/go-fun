package interfaces

import (
	"fmt"
	"strconv"
)

func PrintTypeInfo(v interface{}) {
	switch v2 := v.(type) {
	case int64:
		fmt.Println("Это число " + strconv.FormatInt(v2, 10))
	case string:
		fmt.Println("Это строка " + v2)
	case fmt.Stringer:
		fmt.Println("Это тип, реализующий Stringer, " + v2.String())
	default:
		fmt.Println("Неизвестный тип")
	}
}
