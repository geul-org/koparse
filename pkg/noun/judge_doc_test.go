//ff:func feature=noun type=judge control=sequence
//ff:what TestJudgeDoc — 1차에 주제로 채택된 명사의 분류가 2차에 다른 명사의 R4b 근거가 된다, 단위 오류는 넘긴다
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestJudgeDoc(t *testing.T) {
	lex, _ := lexicon.Default()
	river := []string{"Q4022"}
	us := []Unit{
		{Form: "한강", Matched: "한강", Cands: []knowledge.Candidate{{QID: "Q1"}}, Infos: []knowledge.EntityInfo{{P31: river}}},
		{Form: "낙동강", Matched: "낙동강", Cands: []knowledge.Candidate{{QID: "Q2"}, {QID: "Q3"}}, Infos: []knowledge.EntityInfo{{P31: river}, {}}},
	}
	rs, err := JudgeDoc(us, DocContext{Topic: "Q1"}, lex, config.Default().Noun)
	if err != nil || rs[0].QID != "Q1" || rs[1].QID != "Q2" || rs[1].Rules[0] != RuleR4B {
		t.Fatal(rs, err)
	}
	if _, err := JudgeDoc([]Unit{{Form: "x", Cands: make([]knowledge.Candidate, 1)}}, DocContext{}, lex, config.Default().Noun); err == nil {
		t.Fatal("want error")
	}
}
