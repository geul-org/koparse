//ff:func feature=stats type=builder control=sequence
//ff:what TestSuffixTokens — 이어진 조사는 J 하나, 서술격·용언화는 V(조사+용언화는 J, V)
package stats

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestSuffixTokens(t *testing.T) {
	w := noun.Word{Chain: []lexicon.Suffix{{Form: "에서", Kind: "particle"}, {Form: "는", Kind: "particle"}}}
	if got := suffixTokens(nil, w, 0, 0); len(got) != 1 || got[0].Text != "에서는" || got[0].Tag != "J" {
		t.Fatal(got)
	}
	w2 := noun.Word{Chain: []lexicon.Suffix{{Form: "으로", Kind: "particle"}, {Form: "한", Kind: "verbal"}}}
	got := suffixTokens(nil, w2, 0, 0)
	if len(got) != 2 || got[0].Tag != "J" || got[1].Text != "한" || got[1].Tag != "V" {
		t.Fatal(got)
	}
}
