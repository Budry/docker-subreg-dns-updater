package utils

import (
	"strings"
)

func SplitKeyValueSlice(in []string) map[string]string {
	env := make(map[string]string)
	for _, entry := range in {
		parts := strings.SplitN(entry, "=", 2)
		if len(parts) != 2 {
			parts = append(parts, "")
		}
		env[parts[0]] = parts[1]
	}
	return env
}