//ff:func feature=noun type=judge control=sequence
//ff:what TestJudgeUnit — 후보 없음은 버림, 주제 일치는 채택, 후보·요약 수가 다르면 오류
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestJudgeUnit(t *testing.T) {
	lex, _ := lexicon.Default()
	cfg := config.Default().Noun
	if r, _ := judgeUnit(Unit{Form: "모름"}, DocContext{}, lex, cfg); r.Status != StatusDropped || r.Reason != ReasonNoCandidate {
		t.Fatal(r)
	}
	u := Unit{Form: "김치", Matched: "김치", Cands: []knowledge.Candidate{{QID: "Q1"}, {QID: "Q2"}}, Infos: []knowledge.EntityInfo{{}, {}}}
	r, err := judgeUnit(u, DocContext{Topic: "Q1"}, lex, cfg)
	if err != nil || r.Status != StatusAdopted || r.QID != "Q1" {
		t.Fatal(r, err)
	}
	if _, err := judgeUnit(Unit{Form: "x", Cands: u.Cands}, DocContext{}, lex, cfg); err == nil {
		t.Fatal("want error for missing summaries")
	}
}
