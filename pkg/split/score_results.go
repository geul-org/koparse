//ff:func feature=split type=judge control=iteration dimension=1
//ff:what scoreResults — toulmin 결과를 후보 번호로 모은다(결과가 없는 후보는 verdict −1)
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// scoreResults maps warrant results back to candidates by their evidence index.
func scoreResults(cands []Candidate, results []toulmin.EvalResult) []Scored {
	out := make([]Scored, len(cands))
	for i, c := range cands {
		out[i] = Scored{Candidate: c, Verdict: -1}
	}
	for _, r := range results {
		i := r.Evidence.(int)
		out[i].Verdict, out[i].Applicability = r.Verdict, r.Applicability
	}
	return out
}
