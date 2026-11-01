//ff:func feature=noun type=judge control=sequence
//ff:what TestCombine — 후보별 noisy-OR, 최대 적용도, 끊긴 warrant는 규칙 목록에서 뺀다
package noun

import (
	"math"
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestCombine(t *testing.T) {
	res := []toulmin.EvalResult{
		{Verdict: 0, Applicability: 1, Evidence: hit{Cand: 0, Rule: RuleR1}},
		{Verdict: 0, Applicability: 0.5, Evidence: hit{Cand: 0, Rule: RuleR3}},
		{Verdict: -1, Applicability: 0, Evidence: hit{Cand: 1, Rule: RuleR2}},
	}
	s := combine(3, res)
	if math.Abs(s[0].Miss-0.25) > 1e-9 || s[0].Applicability != 1 || len(s[0].Rules) != 2 || s[1].Miss != 1 || len(s[1].Rules) != 0 || s[2].Miss != 1 {
		t.Fatal(s)
	}
}
