//ff:func feature=eval type=judge control=sequence
//ff:what TestCountRecall — 맞는 QID가 없는 정답은 빼고, 같은 문장 채택만 센다
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestCountRecall(t *testing.T) {
	outs := map[string]analyze.SentenceOut{"1.1": {Nouns: []noun.Result{{Form: "물", QID: "Q1", Status: noun.StatusAdopted}}}}
	gold := []GoldNoun{{Key: "1.1", Noun: "물", Good: []string{"Q1"}}, {Key: "1.1", Noun: "물", Good: []string{"Q1"}}, {Key: "1.1", Noun: "자리"}, {Key: "1.2", Noun: "불", Good: []string{"Q2"}}}
	s := NounScore{}
	countRecall(&s, outs, gold)
	if s.GoldGood != 3 || s.Recalled != 1 {
		t.Fatalf("%+v", s)
	}
}
