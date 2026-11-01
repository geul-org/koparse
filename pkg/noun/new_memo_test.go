//ff:func feature=noun type=builder control=sequence
//ff:what TestNewMemo — 빈 캐시를 가진다
package noun

import "testing"

func TestNewMemo(t *testing.T) {
	m := NewMemo(nil, nil)
	if m.cands == nil || m.ents == nil {
		t.Fatal(m)
	}
}
