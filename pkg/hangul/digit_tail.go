//ff:func feature=hangul type=helper control=selection
//ff:what digitTail — 숫자 끝자리의 한자어 읽기(영·일·이…구)로 끝소리를 정한다
package hangul

// digitTail returns the ending sound of a number by its last digit read in
// Sino-Korean: 0 영(십·백·천·만도 받침), 1 일·7 칠·8 팔 ㄹ, 2 이·4 사·5 오·9 구 모음.
func digitTail(r rune) Tail {
	switch r {
	case '1', '7', '8':
		return TailRieul
	case '2', '4', '5', '9':
		return TailVowel
	default:
		return TailConsonant
	}
}
