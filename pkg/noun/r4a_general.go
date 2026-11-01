//ff:func feature=noun type=rule control=sequence
//ff:what r4aGeneral — 일반 R4a: 고유 후보(인명·이름·글자·작품·연호 등)를 빼고 남은 일반 후보가 하나면 지지
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// r4aGeneral supports the only general candidate among several.
func r4aGeneral(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f, i := factsOf(ctx, specs)
	return verdictIf(f.cands[i].General), hit{Cand: i, Rule: RuleR4A}
}
