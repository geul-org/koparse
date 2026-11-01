//ff:func feature=eval type=judge control=sequence
//ff:what TestScoreNouns — 모든 문장의 명사를 모으고 재현까지 센다
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestScoreNouns(t *testing.T) {
	outs := map[string]analyze.SentenceOut{
		"1.1": {Nouns: []noun.Result{{Form: "물", QID: "Q1", Status: noun.StatusAdopted}}},
		"1.2": {Nouns: []noun.Result{{Form: "불", Status: noun.StatusDropped, Reason: noun.ReasonBelow}}},
	}
	s := ScoreNouns(outs, []GoldNoun{{Key: "1.1", Noun: "물", Good: []string{"Q1"}}}, nil)
	if s.Nouns != 2 || s.Adopted != 1 || s.Judged[Correct] != 1 || s.Recalled != 1 || s.DroppedReason[noun.ReasonBelow] != 1 {
		t.Fatalf("%+v", s)
	}
}
