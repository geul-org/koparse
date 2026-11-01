//ff:func feature=noun type=helper control=sequence
//ff:what TestBestEnding — 어간 전체는 조각이 아니고, 아는 조각 칸을 고르며, 오류는 넘긴다
package noun

import (
	"context"
	"testing"
)

func TestBestEnding(t *testing.T) {
	m := newTestMemo(t, "오류")
	rs := []rune("김치땅")
	st := make([]segState, len(rs)+1)
	st[0].ok = true
	s2, _ := m.bestEnding(context.Background(), rs, st, 2, false)
	if !s2.ok || s2.prev != 0 {
		t.Fatal(s2)
	}
	st[2] = s2
	s3, _ := m.bestEnding(context.Background(), rs, st, 3, false)
	if !s3.ok || s3.prev != 2 || s3.pieces != 2 {
		t.Fatal(s3)
	}
	bad := []rune("오류")
	st2 := make([]segState, 3)
	st2[0].ok = true
	if _, err := m.bestEnding(context.Background(), append(bad, '들'), append(st2, segState{}), 2, false); err == nil {
		t.Fatal("want error")
	}
}
