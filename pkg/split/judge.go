//ff:func feature=split type=judge control=sequence
//ff:what judge — 그래프를 평가해 후보 순위를 매기고 1위가 문턱 이상·2위와 여유값 이상 차이면 채택, 아니면 미정
package split

import (
	"sort"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// judge evaluates the split graph over f and decides. Exclusive choice is made
// outside the engine (current TAN rejects cycles, so candidates do not attack
// each other).
func judge(f *facts) (Decision, error) {
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, f)
	results, err := buildGraph(len(f.cands), f.cfg).Evaluate(ctx)
	if err != nil {
		return Decision{}, err
	}
	ranked := scoreResults(f.cands, results)
	sort.SliceStable(ranked, func(i, j int) bool { return rankLess(ranked[i], ranked[j]) })
	d := Decision{Ranked: ranked, Status: StatusAdopted}
	if ranked[0].Verdict < f.cfg.Adopt || (len(ranked) > 1 && ranked[0].Verdict-ranked[1].Verdict < f.cfg.Margin) {
		d.Status = StatusUndecided
	}
	return d, nil
}
