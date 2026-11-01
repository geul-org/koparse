//ff:func feature=lexicon type=rule control=iteration dimension=1
//ff:what TestChainValid — 자리 오름차순, 자리 4 이상 뒤 금지, @allow 예외
package lexicon

import "testing"

func TestChainValid(t *testing.T) {
	l := &Lexicon{Allow: map[string]bool{"만\t으로": true}}
	s := func(form string, slot int) Suffix { return Suffix{Form: form, Slot: slot} }
	cases := []struct {
		chain []Suffix
		want  bool
	}{
		{nil, true},
		{[]Suffix{s("들", 1), s("에게", 2), s("는", 5)}, true},
		{[]Suffix{s("은", 5), s("들", 1)}, false},
		{[]Suffix{s("을", 4), s("이다", 9)}, false},
		{[]Suffix{s("에", 2), s("에서", 2)}, false},
		{[]Suffix{s("만", 3), s("으로", 2)}, true},
	}
	for _, c := range cases {
		if got := l.ChainValid(c.chain); got != c.want {
			t.Errorf("%v = %v, want %v", c.chain, got, c.want)
		}
	}
}
