package decorator

import (
	"fmt"
	"strings"
)

func GetActionName(handler any) string {
	return strings.Split(fmt.Sprintf("%T", handler), ".")[1]
}
