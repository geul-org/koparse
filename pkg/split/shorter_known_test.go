//ff:func feature=split type=rule control=sequence
//ff:what TestShorterKnown — 물로+부터는 아는 물+로부터에 공격받고, 아는 어간·사슬이 틀린 짧은 후보는 공격하지 않는다
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestShorterKnown(t *testing.T) {
	robuteo := lexicon.Suffix{Form: "로부터", Cond: "VL", Slot: 2}
	buteo := lexicon.Suffix{Form: "부터", Cond: "*", Slot: 2}
	ga := lexicon.Suffix{Form: "가", Cond: "V", Slot: 4}
	e := lexicon.Suffix{Form: "에", Cond: "*", Slot: 2}
	f := &facts{lex: &lexicon.Lexicon{},
		cands: []Candidate{
			{Stem: "물로부터"}, {Stem: "물로", Suffixes: []lexicon.Suffix{buteo}}, {Stem: "물", Suffixes: []lexicon.Suffix{robuteo}},
			{Stem: "방법가", Suffixes: []lexicon.Suffix{e}}, {Stem: "방법", Suffixes: []lexicon.Suffix{ga, e}},
		},
		stemCount: []int{0, 0, 3, 0, 2}}
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, f)
	hit, ev := shorterKnown(ctx, toulmin.Specs{candSpec{Idx: 1}})
	known, _ := shorterKnown(ctx, toulmin.Specs{candSpec{Idx: 2}})
	illFormed, _ := shorterKnown(ctx, toulmin.Specs{candSpec{Idx: 3}})
	if hit != 1 || ev != "물" || known != -1 || illFormed != -1 {
		t.Fatal(hit, ev, known, illFormed)
	}
}
