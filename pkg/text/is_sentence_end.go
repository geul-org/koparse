//ff:func feature=text type=helper control=selection
//ff:what isSentenceEnd — i번째 글자가 문장 끝인지(줄바꿈, 또는 공백·끝이 뒤따르는 . ? !)
package text

import "unicode"

// isSentenceEnd reports whether rs[i] ends a sentence: a newline, or . ? !
// followed by whitespace or the end (so 605.25 does not split).
func isSentenceEnd(rs []rune, i int) bool {
	switch rs[i] {
	case '\n':
		return true
	case '.', '?', '!':
		return i+1 == len(rs) || unicode.IsSpace(rs[i+1])
	default:
		return false
	}
}
