//ff:func feature=hangul type=helper control=iteration dimension=1
//ff:what IsAllHangul — 빈 문자열이 아니고 모든 글자가 한글 음절인지
package hangul

// IsAllHangul reports whether s is non-empty and every rune is a Hangul syllable.
func IsAllHangul(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if !IsSyllable(r) {
			return false
		}
	}
	return true
}
