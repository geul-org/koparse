//ff:func feature=noun type=judge control=sequence
//ff:what TestJudgeAll — 단위마다 결과 하나, 단위 오류는 넘긴다
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestJudgeAll(t *testing.T) {
	lex, _ := lexicon.Default()
	rs, err := judgeAll([]Unit{{Form: "가"}, {Form: "나"}}, DocContext{}, lex, config.Default().Noun)
	if err != nil || len(rs) != 2 || rs[1].Form != "나" {
		t.Fatal(rs, err)
	}
	if _, err := judgeAll([]Unit{{Form: "x", Cands: make([]knowledge.Candidate, 1)}}, DocContext{}, lex, config.Default().Noun); err == nil {
		t.Fatal("want error")
	}
}
