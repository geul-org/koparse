//ff:func feature=eval type=judge control=sequence
//ff:what TestRecallHit — 같은 형태·맞는 QID의 채택 명사를 한 번만 쓴다
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/noun"
)

func TestRecallHit(t *testing.T) {
	ns := []noun.Result{{Form: "물", QID: "Q1", Status: noun.StatusAdopted}, {Form: "물", QID: "Q1", Status: noun.StatusDropped}}
	used := make([]bool, 2)
	g := GoldNoun{Noun: "물", Good: []string{"Q1"}}
	if !recallHit(ns, used, g) || recallHit(ns, used, g) || !used[0] {
		t.Fatal(used)
	}
}
