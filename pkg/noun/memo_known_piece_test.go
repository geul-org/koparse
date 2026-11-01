//ff:func feature=noun type=helper control=sequence
//ff:what TestMemoKnownPiece — 두 음절은 후보만 있으면, 한 음절은 고유 분류 아닌 후보가 있어야, 오류는 넘긴다
package noun

import (
	"context"
	"testing"
)

func TestMemoKnownPiece(t *testing.T) {
	m := newTestMemo(t, "오")
	ctx := context.Background()
	two, _ := m.knownPiece(ctx, "연안")
	water, _ := m.knownPiece(ctx, "물")
	syllable, _ := m.knownPiece(ctx, "속")
	none, _ := m.knownPiece(ctx, "없음")
	if !two || !water || syllable || none {
		t.Fatal(two, water, syllable, none)
	}
	if _, err := m.knownPiece(ctx, "오"); err == nil {
		t.Fatal("want error")
	}
}
