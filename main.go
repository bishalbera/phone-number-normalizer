package main

import (
	"strings"
	"unicode"
)

func main() {
	
}

func normalizeNumber(phone string) string {
	var normalizedNumber strings.Builder

	for _, char := range phone {
		if unicode.IsDigit(char) {
			normalizedNumber.WriteRune(char)
		}
	}
	return normalizedNumber.String()
}
