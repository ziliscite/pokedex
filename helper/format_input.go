package helper

import (
	"strings"
)

func FormatInput(inp string) (cleaned []string) {
	cleaned = make([]string, 2)
	cleaned = strings.Split(inp, " ")
	for i := 0; i < len(cleaned); i++ {
		cleaned[i] = strings.ToLower(cleaned[i])
	}

	if len(cleaned) == 1 {
		cleaned = append(cleaned, "")
	}
	return
}
