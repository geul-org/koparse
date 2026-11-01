//ff:func feature=noun type=judge control=selection
//ff:what decide — 채택·버림·미정: 예외로 끊겼으면 성질·관용, 1위가 문턱 미만이면 버림, 2위와 차이가 여유값 미만이면 미정
package noun

import "github.com/geul-org/koparse/pkg/config"

// decide turns candidate scores into a Result.
func decide(u Unit, f nounFacts, scores []candScore, cfg config.Noun) Result {
	best, second := rankScores(scores)
	b := scores[best]
	r := Result{Form: u.Form, Best: f.cands[best].QID, Verdict: 1 - 2*b.Miss, Applicability: b.Applicability,
		Rules: b.Rules, Candidates: len(scores), Eojeol: [2]int{u.First, u.Last}, Piece: u.Piece}
	if u.Matched != u.Form {
		r.Rules = append(r.Rules, RuleR5)
	}
	gap := 2.0
	if second >= 0 {
		gap = 2 * (scores[second].Miss - b.Miss)
	}
	switch {
	case b.Applicability == 0 && f.property:
		r.Status, r.Reason = StatusDropped, ReasonProperty
	case b.Applicability == 0 && f.idiom:
		r.Status, r.Reason = StatusDropped, ReasonIdiom
	case r.Verdict < cfg.Adopt:
		r.Status, r.Reason = StatusDropped, ReasonBelow
	case gap < cfg.Margin:
		r.Status, r.Reason = StatusUndecided, ReasonAmbiguous
	default:
		r.Status, r.QID = StatusAdopted, r.Best
	}
	return r
}
