//ff:func feature=noun type=rule control=sequence
//ff:what TestR4aGeneral — R4a: 일반 후보 유일 표시가 있는 후보만 +1, 근거는 (후보, 규칙)
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestR4aGeneral(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{General: true}}})
	v, ev := r4aGeneral(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{}}})
	w, _ := r4aGeneral(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	if v != 1 || w != -1 || ev != (hit{Cand: 0, Rule: RuleR4A}) {
		t.Fatal(v, w, ev)
	}
}
