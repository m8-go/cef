package cef

import "strings"

var headerFieldReplacer = strings.NewReplacer(
	"\n", `\n`,
	`\`, `\\`,
	`|`, `\|`,
)

func escapeHeaderField(str string) string { return headerFieldReplacer.Replace(str) }
