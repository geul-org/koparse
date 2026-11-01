//ff:func feature=noun type=helper control=sequence
//ff:what TestMemoEntity — 요약을 캐시하고, 없는 개체는 빈 요약, 그 밖의 오류는 넘긴다
package noun

import (
	"context"
	"testing"
)

func TestMemoEntity(t *testing.T) {
	m := newTestMemo(t, "")
	e, err := m.Entity(context.Background(), "Q1")
	if err != nil || e.Labels["ko"] != "김치" || m.ents["Q1"].QID != "Q1" {
		t.Fatal(e, err)
	}
	if e, err := m.Entity(context.Background(), "Q1"); err != nil || e.QID != "Q1" {
		t.Fatal(e, err)
	}
	if e, err := m.Entity(context.Background(), "Q999"); err != nil || e.QID != "Q999" {
		t.Fatal(e, err)
	}
	if _, err := m.Entity(context.Background(), "QERR"); err == nil {
		t.Fatal("want error")
	}
}
