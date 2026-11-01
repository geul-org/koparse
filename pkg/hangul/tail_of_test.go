//ff:func feature=hangul type=helper control=iteration dimension=1
//ff:what TestTailOf — 빈 문자열·한글·숫자·영문·기호의 끝소리
package hangul

import "testing"

func TestTailOf(t *testing.T) {
	cases := map[string]Tail{"": TailUnknown, "학교": TailVowel, "서울": TailRieul, "한국": TailConsonant,
		"km2": TailVowel, "60": TailConsonant, "GPU": TailUnknown, "0.61%": TailUnknown}
	for s, want := range cases {
		if got := TailOf(s); got != want {
			t.Errorf("TailOf(%q) = %v, want %v", s, got, want)
		}
	}
}
