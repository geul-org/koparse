//ff:func feature=noun type=builder control=sequence
//ff:what TestRunUnits — 구간이 있으면 건너뛰고, 없으면 어절 하나씩, 구간·어절 조회 오류는 넘긴다
package noun

import (
	"context"
	"testing"
)

func TestRunUnits(t *testing.T) {
	m := newTestMemo(t, "오류")
	w := []Word{{Surface: "강", Stem: "강"}, {Surface: "하구", Stem: "하구"}, {Surface: "김치", Stem: "김치"}}
	us, err := m.runUnits(context.Background(), w, 0, 2)
	if err != nil || len(us) != 2 || us[0].Form != "강 하구" || us[0].Last != 1 || us[1].Form != "김치" {
		t.Fatal(us, err)
	}
	bad := []Word{{Surface: "오류", Stem: "오류"}, {Surface: "김치", Stem: "김치"}}
	if _, err := newTestMemo(t, "오류 김치").runUnits(context.Background(), bad, 0, 1); err == nil {
		t.Fatal("want span error")
	}
	if _, err := m.runUnits(context.Background(), bad, 0, 0); err == nil {
		t.Fatal("want single error")
	}
}
