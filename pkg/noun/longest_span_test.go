//ff:func feature=noun type=builder control=sequence
//ff:what TestLongestSpan — 띄운 형태로 찾은 구간, 붙인 형태 시도, 없음, 오류
package noun

import (
	"context"
	"testing"
)

func TestLongestSpan(t *testing.T) {
	m := newTestMemo(t, "오류 김치")
	w := []Word{{Stem: "발효"}, {Stem: "식품"}, {Stem: "김치"}}
	u, ok, err := m.longestSpan(context.Background(), w, 0, 2)
	if err != nil || !ok || u.Form != "발효 식품" || u.First != 0 || u.Last != 1 {
		t.Fatal(u, ok, err)
	}
	if _, ok, _ := m.longestSpan(context.Background(), w, 1, 2); ok {
		t.Fatal("식품 김치 is not a span")
	}
	e := []Word{{Stem: "오류"}, {Stem: "김치"}}
	if _, _, err := m.longestSpan(context.Background(), e, 0, 1); err == nil {
		t.Fatal("want error")
	}
}
