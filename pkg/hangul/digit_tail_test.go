//ff:func feature=hangul type=helper control=iteration dimension=1
//ff:what TestDigitTail — 1·7·8 ㄹ, 2·4·5·9 모음, 0·3·6 받침
package hangul

import "testing"

func TestDigitTail(t *testing.T) {
	want := map[rune]Tail{'0': TailConsonant, '1': TailRieul, '2': TailVowel, '3': TailConsonant, '4': TailVowel,
		'5': TailVowel, '6': TailConsonant, '7': TailRieul, '8': TailRieul, '9': TailVowel}
	for r, w := range want {
		if got := digitTail(r); got != w {
			t.Errorf("digitTail(%q) = %v, want %v", r, got, w)
		}
	}
}
