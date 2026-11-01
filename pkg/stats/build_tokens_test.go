//ff:func feature=stats type=builder control=iteration dimension=1
//ff:what TestBuildTokens — 지은 문장 "강 하구는 넓고, 물이 맑다": N J V | N J V, 쉼표 뒤 절 번호 증가
package stats

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestBuildTokens(t *testing.T) {
	eun := []lexicon.Suffix{{Form: "는", Kind: "particle"}}
	i := []lexicon.Suffix{{Form: "이", Kind: "particle"}}
	words := []noun.Word{{Surface: "강", Stem: "강"}, {Surface: "하구는", Stem: "하구", Chain: eun}, {Surface: "넓고", Stem: "넓고", Break: true},
		{Surface: "물이", Stem: "물", Chain: i}, {Surface: "맑다", Stem: "맑다"}}
	units := []noun.Unit{{Form: "강 하구", First: 0, Last: 1}, {Form: "물", First: 3, Last: 3}}
	toks := buildTokens(words, units)
	tags := ""
	for _, t := range toks {
		tags += t.Tag
	}
	if tags != "NJVNJV" || toks[3].Clause != 1 || toks[0].Clause != 0 {
		t.Fatal(toks)
	}
}
