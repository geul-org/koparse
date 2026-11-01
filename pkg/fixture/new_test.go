//ff:func feature=fixture type=builder control=sequence
//ff:what TestNew — nil 표도 빈 표로 쓸 수 있다
package fixture

import (
	"context"
	"testing"
)

func TestNew(t *testing.T) {
	f := New(nil, nil, nil)
	c, err := f.Candidates(context.Background(), "김치", "ko")
	s, _ := f.Statements(context.Background(), "Q1", nil)
	if err != nil || c != nil || len(s) != 0 || f.misses == nil {
		t.Fatal(c, s, err)
	}
}
