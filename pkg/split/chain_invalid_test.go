//ff:func feature=split type=rule control=sequence
//ff:what TestChainInvalid — 은+들은 +1, 들+은은 −1
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestChainInvalid(t *testing.T) {
	deul, eun := lexicon.Suffix{Form: "들", Slot: 1}, lexicon.Suffix{Form: "은", Slot: 5}
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, &facts{lex: &lexicon.Lexicon{}, cands: []Candidate{{}, {Stem: "선수", Suffixes: []lexicon.Suffix{eun, deul}}, {Stem: "선수", Suffixes: []lexicon.Suffix{deul, eun}}}})
	bad, _ := chainInvalid(ctx, toulmin.Specs{candSpec{Idx: 1}})
	ok, _ := chainInvalid(ctx, toulmin.Specs{candSpec{Idx: 2}})
	if bad != 1 || ok != -1 {
		t.Fatal(bad, ok)
	}
}
