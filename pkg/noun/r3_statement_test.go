//ff:func feature=noun type=rule control=sequence
//ff:what TestR3Statement — R3: 주제 진술 값인 후보만 +1, 근거는 (후보, 규칙)
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestR3Statement(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{Statement: true}}})
	v, ev := r3Statement(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{}}})
	w, _ := r3Statement(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	if v != 1 || w != -1 || ev != (hit{Cand: 0, Rule: RuleR3}) {
		t.Fatal(v, w, ev)
	}
}
