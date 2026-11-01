//ff:func feature=stats type=builder control=sequence
//ff:what TestRecords — 지은 문장 "물이 맑지 않다": 명사마다 레코드, 조사·지배어·부정·쌍
package stats

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestRecords(t *testing.T) {
	words := []noun.Word{{Surface: "물이", Stem: "물", Chain: []lexicon.Suffix{{Form: "이", Kind: "particle"}}}, {Surface: "맑지", Stem: "맑지"}, {Surface: "않다", Stem: "않다"}}
	units := []noun.Unit{{Form: "물", First: 0, Last: 0}}
	rs := Records(words, units, []noun.Result{{QID: "Q283", Status: noun.StatusAdopted}})
	if len(rs) != 1 || rs[0].Noun != "물" || rs[0].QID != "Q283" || rs[0].Josa != "이" || rs[0].Governor != "맑지" || !rs[0].Negated || len(rs[0].Pairs) != 3 {
		t.Fatalf("%+v", rs)
	}
	if len(Records(words, nil, nil)) != 0 {
		t.Fatal("no units, no records")
	}
}
