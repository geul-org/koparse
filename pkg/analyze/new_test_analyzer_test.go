//ff:func feature=analyze type=helper control=sequence
//ff:what newTestAnalyzer — 시험용 처리기: 저장소 fixture를 감싸 failOn 단계(candidates·entity·statements)에서 오류
package analyze

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/fixture"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// newTestAnalyzer returns an Analyzer over the repository fixture whose
// failOn method ("candidates", "entity", "statements") fails.
func newTestAnalyzer(t *testing.T, failOn string) *Analyzer {
	fx, err := fixture.Load("../../testdata/fixture.json")
	if err != nil {
		t.Fatal(err)
	}
	lex, _ := lexicon.Default()
	boom := errors.New("boom")
	k := knowledge.Funcs{
		CandidatesFn: func(ctx context.Context, s, l string) ([]knowledge.Candidate, error) {
			if failOn == "candidates" {
				return nil, boom
			}
			return fx.Candidates(ctx, s, l)
		},
		EntityFn: func(ctx context.Context, q string) (knowledge.EntityInfo, error) {
			if failOn == "entity" {
				return knowledge.EntityInfo{}, boom
			}
			return fx.Entity(ctx, q)
		},
		StatementsFn: func(ctx context.Context, q string, p []string) (map[string][]string, error) {
			if failOn == "statements" {
				return nil, boom
			}
			return fx.Statements(ctx, q, p)
		},
	}
	return New(k, lex, config.Default())
}
