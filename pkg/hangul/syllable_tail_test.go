//ff:func feature=hangul type=helper control=iteration dimension=1
//ff:what TestSyllableTail — 사(모음)·물(ㄹ)·국(받침)
package hangul

import "testing"

func TestSyllableTail(t *testing.T) {
	for r, want := range map[rune]Tail{'사': TailVowel, '물': TailRieul, '국': TailConsonant} {
		if got := syllableTail(r); got != want {
			t.Errorf("syllableTail(%q) = %v, want %v", r, got, want)
		}
	}
}
