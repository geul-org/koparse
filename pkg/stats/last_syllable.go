//ff:func feature=stats type=helper control=sequence
//ff:what lastSyllable — 마지막 글자(없으면 0)
package stats

import "unicode/utf8"

// lastSyllable returns the last rune of s (0 for "").
func lastSyllable(s string) rune {
	if s == "" {
		return 0
	}
	r, _ := utf8.DecodeLastRuneInString(s)
	return r
}
