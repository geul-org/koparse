//ff:func feature=noun type=helper control=sequence
//ff:what TestMemoLookup — 위키 내부 문서는 빼고 캐시하며, 후보·개체 오류는 넘긴다
package noun

import (
	"context"
	"testing"
)

func TestMemoLookup(t *testing.T) {
	m := newTestMemo(t, "오류")
	r, err := m.lookup(context.Background(), "한강")
	if err != nil || len(r.cands) != 2 || r.infos[0].P31[0] != "Q4022" {
		t.Fatal(r, err)
	}
	if r, _ := m.lookup(context.Background(), "위키"); len(r.cands) != 0 {
		t.Fatal("internal page kept", r)
	}
	if r, _ := m.lookup(context.Background(), "한강"); len(r.cands) != 2 {
		t.Fatal("cache", r)
	}
	if _, err := m.lookup(context.Background(), "오류"); err == nil {
		t.Fatal("want candidates error")
	}
	if _, err := m.lookup(context.Background(), "고장"); err == nil {
		t.Fatal("want entity error")
	}
}
