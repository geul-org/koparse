//ff:func feature=noun type=rule control=sequence
//ff:what r2Topic — 일반 R2: 후보가 문서 주제 개체면 지지
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// r2Topic supports the candidate that is the document topic.
func r2Topic(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f, i := factsOf(ctx, specs)
	return verdictIf(f.cands[i].Topic), hit{Cand: i, Rule: RuleR2}
}
