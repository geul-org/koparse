//ff:func feature=eval type=helper control=sequence
//ff:what newTestAnalyzer — 시험용: 저장소 fixture 위 처리기(fail이면 후보 조회가 실패)
package eval

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/fixture"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// newTestAnalyzer returns an Analyzer over the repository fixture.
func newTestAnalyzer(t *testing.T, fail bool) *analyze.Analyzer {
	fx, err := fixture.Load("../../testdata/fixture.json")
	if err != nil {
		t.Fatal(err)
	}
	lex, _ := lexicon.Default()
	k := knowledge.Funcs{
		CandidatesFn: func(ctx context.Context, s, l string) ([]knowledge.Candidate, error) {
			if fail {
				return nil, errors.New("boom")
			}
			return fx.Candidates(ctx, s, l)
		},
		EntityFn: fx.Entity, StatementsFn: fx.Statements,
	}
	return analyze.New(k, lex, config.Default())
}
