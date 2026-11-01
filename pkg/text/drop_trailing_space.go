//ff:func feature=text type=helper control=sequence
//ff:what dropTrailingSpace — 끝 공백 한 글자를 뗀다
package text

import "unicode"

// dropTrailingSpace removes one trailing whitespace rune from b.
func dropTrailingSpace(b []rune) []rune {
	if len(b) > 0 && unicode.IsSpace(b[len(b)-1]) {
		return b[:len(b)-1]
	}
	return b
}
