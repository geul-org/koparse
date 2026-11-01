//ff:func feature=text type=helper control=sequence
//ff:what isTrimRune — 어절 앞뒤에서 떼는 문장부호인지(% - 는 남긴다)
package text

import "unicode"

// isTrimRune reports whether r is punctuation stripped from token edges.
// '%' and '-' stay attached (0.61%에, -特別市).
func isTrimRune(r rune) bool {
	return unicode.IsPunct(r) && r != '%' && r != '-'
}
