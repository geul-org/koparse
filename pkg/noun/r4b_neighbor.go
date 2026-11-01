//ff:func feature=noun type=rule control=sequence
//ff:what r4bNeighbor — 일반 R4b: 일반 후보 중 문서에서 채택된 이웃 명사와 분류(P31·P279)를 나누는 것이 하나면 지지
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// r4bNeighbor supports the only general candidate that shares a class with
// nouns adopted elsewhere in the document.
func r4bNeighbor(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f, i := factsOf(ctx, specs)
	return verdictIf(f.cands[i].Neighbor), hit{Cand: i, Rule: RuleR4B}
}
