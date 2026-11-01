//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what TestSegmentCompound — 아는 조각으로 나누고, 접미가 있을 때만 가장자리 접사, 짧은·섞인 어간과 한 음절 조각뿐인 어간은 안 나눈다
package noun

import (
	"context"
	"testing"
)

func TestSegmentCompound(t *testing.T) {
	m := newTestMemo(t, "오류")
	ctx := context.Background()
	p, err := m.segmentCompound(ctx, "김치땅", false)
	if err != nil || len(p) != 2 || p[0].form != "김치" || p[1].form != "땅" {
		t.Fatal(p, err)
	}
	if p, _ := m.segmentCompound(ctx, "쓴김치", true); len(p) != 2 || p[0].known || !p[1].known {
		t.Fatal(p)
	}
	for _, s := range []string{"쓴김치", "땅물", "김", "GPU김치"} {
		if p, _ := m.segmentCompound(ctx, s, false); p != nil {
			t.Errorf("%s: %v", s, p)
		}
	}
	if _, err := m.segmentCompound(ctx, "오류치", false); err == nil {
		t.Fatal("want error")
	}
}
