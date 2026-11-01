//ff:func feature=eval type=judge control=sequence
//ff:what TestAddAdopted — 판정별·규칙별로 세고 목록 줄을 남긴다
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/noun"
)

func TestAddAdopted(t *testing.T) {
	s := NounScore{Judged: map[string]int{}, ByRule: map[string][3]int{}}
	ix := newGoldIndex([]GoldNoun{{Key: "1.1", Noun: "물", Good: []string{"Q1"}}}, nil)
	addAdopted(&s, ix, "1.1", noun.Result{Form: "물", QID: "Q1", Rules: []string{noun.RuleR2, noun.RuleR3}})
	if s.Judged[Correct] != 1 || s.ByRule[noun.RuleR2] != [3]int{1, 0, 0} || s.Rows[0] != "1.1\t물\tQ1\t맞음\tR2 주제,R3 진술" {
		t.Fatalf("%+v", s)
	}
}
