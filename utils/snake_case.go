package utils

import (
	"strings"
	"unicode"
)

func ToSnakeCase(s string) string {
	var builder strings.Builder
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 { // Add underscore before uppercase letters (except the first)
				builder.WriteByte('_')
			}
			builder.WriteRune(unicode.ToLower(r)) // Convert to lowercase
		} else {
			builder.WriteRune(r) // Append as is
		}
	}
	return builder.String()
}
