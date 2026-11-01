//ff:func feature=split type=builder control=sequence
//ff:what TestBuildGraph — 후보 셋이면 warrant 셋(통째 + 분리 둘)이 평가된다
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestBuildGraph(t *testing.T) {
	lex, _ := lexicon.Default()
	cfg := config.Default().Split
	cands := Generate(lex, "책을")[:2]
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, &facts{cands: cands, stemCount: []int{0, 1}, lex: lex, cfg: cfg})
	res, err := buildGraph(len(cands), cfg).Evaluate(ctx)
	if err != nil || len(res) != 2 {
		t.Fatal(res, err)
	}
}
