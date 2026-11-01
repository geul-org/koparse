//ff:func feature=hangul type=helper control=selection
//ff:what TailOf — 문자열 끝 글자의 끝소리(한글 받침, 숫자 읽기)
package hangul

import "unicode/utf8"

// TailOf returns the ending sound of s: Hangul syllables by their final
// consonant, digits by their Sino-Korean reading, everything else unknown.
func TailOf(s string) Tail {
	r, _ := utf8.DecodeLastRuneInString(s)
	switch {
	case r == utf8.RuneError:
		return TailUnknown
	case IsSyllable(r):
		return syllableTail(r)
	case r >= '0' && r <= '9':
		return digitTail(r)
	default:
		return TailUnknown
	}
}
