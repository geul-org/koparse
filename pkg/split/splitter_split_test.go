//ff:func feature=split type=judge control=sequence
//ff:what TestSplitterSplit — 지은 문장 어절: 책을→책+을, 가수(통째 라벨)→통째, 어댑터 오류는 넘긴다
package split

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/text"
)

func TestSplitterSplit(t *testing.T) {
	lex, _ := lexicon.Default()
	known := map[string]bool{"책": true, "가수": true}
	k := knowledge.Funcs{CandidatesFn: func(_ context.Context, s, _ string) ([]knowledge.Candidate, error) {
		if s == "오류" {
			return nil, errors.New("boom")
		}
		if known[s] {
			return []knowledge.Candidate{{QID: "Q1"}}, nil
		}
		return nil, nil
	}}
	s := NewSplitter(k, lex, config.Default().Split)
	d, err := s.Split(context.Background(), text.Eojeol{Surface: "책을"})
	if err != nil || d.Ranked[0].Candidate.Key() != "책|을:particle" {
		t.Fatal(d, err)
	}
	if d, _ := s.Split(context.Background(), text.Eojeol{Surface: "가수"}); d.Ranked[0].Candidate.Key() != "가수" {
		t.Fatal(d.Ranked[0].Candidate.Key())
	}
	if _, err := s.Split(context.Background(), text.Eojeol{Surface: "오류"}); err == nil {
		t.Fatal("want error")
	}
}
