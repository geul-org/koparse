//ff:func feature=eval type=builder control=sequence
//ff:what TestAddSets — 같은 칸에 여러 줄을 합친다
package eval

import "testing"

func TestAddSets(t *testing.T) {
	m := map[string]nounSets{}
	addSets(m, "k", GoldNoun{Good: []string{"Q1"}})
	addSets(m, "k", GoldNoun{Good: []string{"Q2"}, Bad: []string{"Q3"}})
	if len(m["k"].good) != 2 || !m["k"].bad["Q3"] {
		t.Fatal(m)
	}
}
