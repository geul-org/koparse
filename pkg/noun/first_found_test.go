//ff:func feature=noun type=helper control=sequence
//ff:what TestFirstFound — 후보가 있는 첫 형태, 없으면 빈 형태, 오류는 넘긴다
package noun

import (
	"context"
	"testing"
)

func TestFirstFound(t *testing.T) {
	m := newTestMemo(t, "오류")
	ctx := context.Background()
	f, r, err := m.firstFound(ctx, []string{"없음", "발효 식품", "김치"})
	if err != nil || f != "발효 식품" || r.cands[0].QID != "Q4" {
		t.Fatal(f, r, err)
	}
	if f, _, _ := m.firstFound(ctx, []string{"없음"}); f != "" {
		t.Fatal(f)
	}
	if _, _, err := m.firstFound(ctx, []string{"오류"}); err == nil {
		t.Fatal("want error")
	}
}
