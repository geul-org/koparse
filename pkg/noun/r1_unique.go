//ff:func feature=noun type=rule control=sequence
//ff:what r1Unique — 일반 R1: 후보가 하나뿐이면 지지(단독으로는 문턱을 넘지 못하게 Qualifier를 낮게)
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// r1Unique supports the only candidate.
func r1Unique(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f, i := factsOf(ctx, specs)
	return verdictIf(len(f.cands) == 1), hit{Cand: i, Rule: RuleR1}
}
