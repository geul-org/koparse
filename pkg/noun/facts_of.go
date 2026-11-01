//ff:func feature=noun type=helper control=sequence
//ff:what factsOf — toulmin 문맥에서 판정 재료와 규칙의 후보 번호를 꺼낸다
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// factsOf returns the facts in ctx and the candidate index of the rule's spec
// (-1 for a rule without spec).
func factsOf(ctx toulmin.Context, specs toulmin.Specs) (*nounFacts, int) {
	v, _ := ctx.Get(nounFactsKey)
	s, ok := specs.Find("qid").(nounSpec)
	if !ok {
		return v.(*nounFacts), -1
	}
	return v.(*nounFacts), s.Cand
}
