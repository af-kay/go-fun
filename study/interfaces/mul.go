package interfaces

import (
	"fmt"
	"strings"
)

func Mul(a any, b int64) any {
    switch a2 := a.(type) {
    case int64:
        return a2 * b
    case fmt.Stringer:
        var builder strings.Builder

        for b > 0 {
            builder.WriteString(a2.String())
            b--
        }

        return builder.String()
    case string:
        var builder strings.Builder

        for b > 0 {
            builder.WriteString(a2)
            b--
        }

        return builder.String()
    default:
        return nil
    }
}
