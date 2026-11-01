//ff:func feature=noun type=helper control=sequence
//ff:what TestBetterState — 불가능한 칸은 지고, 조각이 적은 쪽, 같으면 접사가 적은 쪽
package noun

import "testing"

func TestBetterState(t *testing.T) {
	none := segState{}
	two := segState{ok: true, pieces: 2}
	three := segState{ok: true, pieces: 3}
	twoAffix := segState{ok: true, pieces: 2, affixes: 1}
	if betterState(none, two) || !betterState(two, none) || !betterState(two, three) || betterState(three, two) || !betterState(two, twoAffix) || betterState(twoAffix, two) {
		t.Fatal("betterState")
	}
}
