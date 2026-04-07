// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: [Jonathan Ahubi]
// Squad:  [Channels]

package main

import (
	"fmt"
	"strings"
)

func Cap(s string) string {
	s = strings.ToLower(s)
	return strings.Title(s)

}

func Lower(s string) string {
	return strings.ToUpper(s)
}

func Replace(s string) string {
	s = strings.ReplaceAll(s, "TODO", "ACTION")
	return s
}

func Rep(s string) string {
	s = strings.ReplaceAll(s, " CLASSIFIED", " [REDACTED]")
	return s
}

func Trim(s string) string {
	s = strings.TrimSpace(s)
	str := strings.Fields(s)
	return strings.Join(str, " ")
}

