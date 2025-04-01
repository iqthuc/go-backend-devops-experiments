package utils

import (
	"fmt"
	"strings"
)

func BuildPostgresPlaceholders(n int) string {
	if n == 0 {
		return ""
	}

	var sb strings.Builder
	for i := 1; i <= n; i++ {
		sb.WriteString(fmt.Sprintf("$%d, ", i))
	}
	result := sb.String()
	return result[:len(result)-2]
}
