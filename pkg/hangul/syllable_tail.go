//ff:func feature=hangul type=helper control=selection
//ff:what syllableTail — 한글 음절의 종성으로 끝소리 종류를 정한다
package hangul

// syllableTail classifies a precomposed syllable by its final consonant index.
func syllableTail(r rune) Tail {
	switch (r - 0xAC00) % 28 {
	case 0:
		return TailVowel
	case 8:
		return TailRieul
	default:
		return TailConsonant
	}
}
