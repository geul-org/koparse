//ff:func feature=noun type=rule control=sequence
//ff:what TestR2Topic — R2: 주제 개체인 후보만 +1, 근거는 (후보, 규칙)
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestR2Topic(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{Topic: true}}})
	v, ev := r2Topic(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{}}})
	w, _ := r2Topic(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	if v != 1 || w != -1 || ev != (hit{Cand: 0, Rule: RuleR2}) {
		t.Fatal(v, w, ev)
	}
}
