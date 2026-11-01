//ff:func feature=lexicon type=rule control=iteration dimension=1
//ff:what TestSuffixAccepts — 받침 조건 C·V·VL·CX·*와 알 수 없는 끝소리
package lexicon

import (
	"testing"

	"github.com/geul-org/koparse/pkg/hangul"
)

func TestSuffixAccepts(t *testing.T) {
	cases := []struct {
		cond string
		tail hangul.Tail
		want bool
	}{
		{"C", hangul.TailConsonant, true}, {"C", hangul.TailRieul, true}, {"C", hangul.TailVowel, false},
		{"V", hangul.TailVowel, true}, {"V", hangul.TailRieul, false},
		{"VL", hangul.TailRieul, true}, {"VL", hangul.TailVowel, true}, {"VL", hangul.TailConsonant, false},
		{"CX", hangul.TailConsonant, true}, {"CX", hangul.TailRieul, false},
		{"*", hangul.TailVowel, true}, {"V", hangul.TailUnknown, true},
	}
	for _, c := range cases {
		if got := (Suffix{Cond: c.cond}).Accepts(c.tail); got != c.want {
			t.Errorf("%s after %v = %v, want %v", c.cond, c.tail, got, c.want)
		}
	}
}
