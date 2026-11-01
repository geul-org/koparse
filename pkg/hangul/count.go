//ff:func feature=hangul type=helper control=sequence
//ff:what Count — 글자(룬) 수
package hangul

import "unicode/utf8"

// Count returns the number of runes (syllables for Hangul text) in s.
func Count(s string) int {
	return utf8.RuneCountInString(s)
}
