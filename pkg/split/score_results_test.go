//ff:func feature=split type=judge control=sequence
//ff:what TestScoreResults — 결과는 근거 번호의 후보로, 결과 없는 후보는 −1
package split

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestScoreResults(t *testing.T) {
	got := scoreResults([]Candidate{{Stem: "a"}, {Stem: "b"}}, []toulmin.EvalResult{{Verdict: 0.5, Evidence: 1, Applicability: 0.7}})
	if got[0].Verdict != -1 || got[1].Verdict != 0.5 || got[1].Applicability != 0.7 || got[1].Candidate.Stem != "b" {
		t.Fatal(got)
	}
}
