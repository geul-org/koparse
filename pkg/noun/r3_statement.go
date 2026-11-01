//ff:func feature=noun type=rule control=sequence
//ff:what r3Statement — 일반 R3: 후보가 주제 개체 진술의 값이면 지지
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// r3Statement supports a candidate that is a statement value of the topic.
func r3Statement(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f, i := factsOf(ctx, specs)
	return verdictIf(f.cands[i].Statement), hit{Cand: i, Rule: RuleR3}
}
