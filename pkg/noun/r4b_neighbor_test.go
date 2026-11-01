//ff:func feature=noun type=rule control=sequence
//ff:what TestR4bNeighbor — R4b: 이웃 분야 표시가 있는 후보만 +1, 근거는 (후보, 규칙)
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestR4bNeighbor(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{Neighbor: true}}})
	v, ev := r4bNeighbor(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{}}})
	w, _ := r4bNeighbor(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	if v != 1 || w != -1 || ev != (hit{Cand: 0, Rule: RuleR4B}) {
		t.Fatal(v, w, ev)
	}
}
