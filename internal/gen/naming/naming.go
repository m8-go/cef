package naming

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Pretty(str string) string {
	var b strings.Builder

	parts := Split(str)

	for _, part := range parts {
		s, ok := Rule(part)
		if ok {
			b.WriteString(s)
		} else {
			b.WriteString(part)
		}
	}

	return b.String()
}

func Split(str string) []string {
	if !utf8.ValidString(str) {
		return []string{str}
	}
	var parts []string
	var runes [][]rune
	lastClass := 0
	class := 0
	for _, r := range str {
		if r == ' ' {
			continue
		}

		switch {
		case unicode.IsLower(r):
			class = 1
		case unicode.IsUpper(r):
			class = 2
		case unicode.IsDigit(r):
			class = 3
		default:
			class = 4
		}
		if class == lastClass {
			runes[len(runes)-1] = append(runes[len(runes)-1], r)
		} else {
			runes = append(runes, []rune{r})
		}
		lastClass = class
	}
	for i := 0; i < len(runes)-1; i++ {
		if unicode.IsUpper(runes[i][0]) && unicode.IsLower(runes[i+1][0]) {
			runes[i+1] = append([]rune{runes[i][len(runes[i])-1]}, runes[i+1]...)
			runes[i] = runes[i][:len(runes[i])-1]
		}
	}
	for _, s := range runes {
		if len(s) > 0 {
			parts = append(parts, string(s))
		}
	}
	return parts
}
