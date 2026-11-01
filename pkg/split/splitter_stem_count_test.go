//ff:func feature=split type=helper control=sequence
//ff:what TestSplitterStemCount — 후보 수를 세어 캐시하고, 어댑터 오류는 넘긴다
package split

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
)

func TestSplitterStemCount(t *testing.T) {
	calls := 0
	k := knowledge.Funcs{CandidatesFn: func(_ context.Context, s, _ string) ([]knowledge.Candidate, error) {
		calls++
		if s == "오류" {
			return nil, errors.New("boom")
		}
		return []knowledge.Candidate{{QID: "Q1"}, {QID: "Q2"}}, nil
	}}
	s := NewSplitter(k, nil, config.Split{})
	n1, _ := s.stemCount(context.Background(), "책")
	n2, _ := s.stemCount(context.Background(), "책")
	if n1 != 2 || n2 != 2 || calls != 1 {
		t.Fatal(n1, n2, calls)
	}
	if _, err := s.stemCount(context.Background(), "오류"); err == nil {
		t.Fatal("want error")
	}
}
