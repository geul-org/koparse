//ff:func feature=noun type=rule control=sequence
//ff:what TestR1Unique — R1: 후보가 하나일 때만 +1, 근거는 (후보, 규칙)
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestR1Unique(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{}}})
	v, ev := r1Unique(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{}, {}}})
	w, _ := r1Unique(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	if v != 1 || w != -1 || ev != (hit{Cand: 0, Rule: RuleR1}) {
		t.Fatal(v, w, ev)
	}
}
