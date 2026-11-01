//ff:func feature=hangul type=helper control=sequence
//ff:what IsSyllable — 완성형 한글 음절(가–힣)인지
package hangul

// IsSyllable reports whether r is a precomposed Hangul syllable (U+AC00–U+D7A3).
func IsSyllable(r rune) bool {
	return r >= 0xAC00 && r <= 0xD7A3
}
