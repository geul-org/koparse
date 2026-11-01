//ff:func feature=noun type=helper control=sequence
//ff:what TestNextState — 모르는 조각은 접사로만, 접사는 수를 센다
package noun

import "testing"

func TestNextState(t *testing.T) {
	prev := segState{ok: true, pieces: 1}
	if s := nextState(prev, 2, false, false); s.ok {
		t.Fatal(s)
	}
	if s := nextState(prev, 2, true, false); !s.ok || s.pieces != 2 || s.affixes != 0 || s.prev != 2 || !s.known {
		t.Fatal(s)
	}
	if s := nextState(prev, 2, false, true); !s.ok || s.affixes != 1 || s.known {
		t.Fatal(s)
	}
}
