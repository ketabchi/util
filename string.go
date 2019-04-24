package util

import (
	"strings"
)

func filter(r rune) rune {
	// TODO what to do with nimfasele های
	switch r {
	case 8205, 8207, 8235, 8236, 8238:
		return -1
	}
	return r
}

func Clean(s string) string {
	s = strings.Map(filter, s)
	s = strings.Replace(s, "\u200C ", " ", -1)
	s = strings.Replace(s, " : ", ": ", -1)
	s = strings.Replace(s, " ...", "...", -1)
	// TODO: remove if it was 5 spaces
	s = strings.Replace(s, "  ", " ", -1)
	// TODO: remove if it was 5 zero width spaces
	s = strings.Replace(s, "\u200c\u200c", "\u200c", -1)
	s = strings.TrimSuffix(s, string('\u200C'))
	s = strings.TrimSuffix(s, string('\u202D'))

	return strings.TrimSpace(s)
}
