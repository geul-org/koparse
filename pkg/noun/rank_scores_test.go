//ff:func feature=noun type=judge control=sequence
//ff:what TestRankScores — 1위·2위(같으면 앞 후보), 하나뿐이면 2위 −1
package noun

import "testing"

func TestRankScores(t *testing.T) {
	b, s := rankScores([]candScore{{Miss: 0.5}, {Miss: 0.1}, {Miss: 0.3}, {Miss: 0.3}})
	if b != 1 || s != 2 {
		t.Fatal(b, s)
	}
	b, s = rankScores([]candScore{{Miss: 0.1}, {Miss: 0.5}})
	if b != 0 || s != 1 {
		t.Fatal(b, s)
	}
	b, s = rankScores([]candScore{{Miss: 1}})
	if b != 0 || s != -1 {
		t.Fatal(b, s)
	}
}
