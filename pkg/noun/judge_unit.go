//ff:func feature=noun type=judge control=sequence
//ff:what judgeUnit — 명사 하나: 단위 검사(후보·요약 수) → 판정 재료 → TAN 그래프 평가 → 후보별 합침 → 채택·버림·미정
package noun

import (
	"fmt"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// judgeUnit judges one noun unit.
func judgeUnit(u Unit, dc DocContext, lex *lexicon.Lexicon, cfg config.Noun) (Result, error) {
	if len(u.Infos) != len(u.Cands) {
		return Result{}, fmt.Errorf("noun: unit %q has %d candidates but %d summaries", u.Form, len(u.Cands), len(u.Infos))
	}
	f := buildFacts(u, dc, lex)
	if len(f.cands) == 0 {
		return Result{Form: u.Form, Verdict: -1, Status: StatusDropped, Reason: ReasonNoCandidate, Eojeol: [2]int{u.First, u.Last}, Piece: u.Piece}, nil
	}
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &f)
	results, err := buildNounGraph(len(f.cands), cfg).Evaluate(ctx)
	if err != nil {
		return Result{}, err
	}
	return decide(u, f, combine(len(f.cands), results), cfg), nil
}
