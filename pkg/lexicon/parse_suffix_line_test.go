//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what TestParseSuffixLine — 올바른 줄과 필드 수·종류·자리·조건·무게 오류
package lexicon

import "testing"

func TestParseSuffixLine(t *testing.T) {
	s, err := parseSuffixLine([]string{"을", "particle", "4", "C", "0.9"})
	if err != nil || s != (Suffix{Form: "을", Kind: KindParticle, Slot: 4, Cond: "C", Weight: 0.9}) {
		t.Fatalf("got %+v, %v", s, err)
	}
	bad := [][]string{
		{"을", "particle", "4", "C"},
		{"을", "noun", "4", "C", "0.9"},
		{"을", "particle", "x", "C", "0.9"},
		{"을", "particle", "4", "Q", "0.9"},
		{"을", "particle", "4", "C", "x"},
		{"을", "particle", "4", "C", "1.5"},
		{"을", "particle", "4", "C", "-0.1"},
	}
	for _, f := range bad {
		if _, err := parseSuffixLine(f); err == nil {
			t.Errorf("%v: want error", f)
		}
	}
}
