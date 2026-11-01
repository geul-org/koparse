//ff:func feature=eval type=judge control=sequence
//ff:what TestAddNoun — 상태별로 세고 버림은 사유도 센다
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/noun"
)

func TestAddNoun(t *testing.T) {
	s := NounScore{DroppedReason: map[string]int{}, Judged: map[string]int{}, ByRule: map[string][3]int{}}
	ix := newGoldIndex(nil, nil)
	addNoun(&s, ix, "1.1", noun.Result{Status: noun.StatusAdopted, Form: "물", QID: "Q1"})
	addNoun(&s, ix, "1.1", noun.Result{Status: noun.StatusUndecided})
	addNoun(&s, ix, "1.1", noun.Result{Status: noun.StatusDropped, Reason: noun.ReasonProperty})
	if s.Nouns != 3 || s.Adopted != 1 || s.Undecided != 1 || s.Dropped != 1 || s.DroppedReason[noun.ReasonProperty] != 1 || s.Judged[Unjudged] != 1 {
		t.Fatalf("%+v", s)
	}
}
