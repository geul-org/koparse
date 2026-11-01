//ff:func feature=noun type=judge control=iteration dimension=1
//ff:what combine — toulmin 결과를 후보별로 합친다: raw = (verdict+1)/2의 noisy-OR, 적용도 최대, 지지가 남은 규칙
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// combine folds warrant results into per-candidate scores.
func combine(n int, results []toulmin.EvalResult) []candScore {
	out := make([]candScore, n)
	for i := range out {
		out[i].Miss = 1
	}
	for _, r := range results {
		h := r.Evidence.(hit)
		s := &out[h.Cand]
		s.Miss *= 1 - (r.Verdict+1)/2
		s.Applicability = max(s.Applicability, r.Applicability)
		if r.Verdict > -1 {
			s.Rules = append(s.Rules, h.Rule)
		}
	}
	return out
}
