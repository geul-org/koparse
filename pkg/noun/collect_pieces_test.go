//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what TestCollectPieces — 끝 칸이 불가·조각 하나·접사뿐이면 nil, 아니면 거슬러 모은다, 오류는 넘긴다
package noun

import (
	"context"
	"testing"
)

func TestCollectPieces(t *testing.T) {
	m := newTestMemo(t, "오류")
	ctx := context.Background()
	rs := []rune("김치땅")
	st := []segState{{ok: true}, {}, {ok: true, pieces: 1, known: true}, {ok: true, pieces: 2, prev: 2, known: true}}
	p, err := m.collectPieces(ctx, rs, st)
	if err != nil || len(p) != 2 || p[0].form != "김치" || p[1].form != "땅" || len(p[1].r.cands) != 1 {
		t.Fatal(p, err)
	}
	for _, end := range []segState{{}, {ok: true, pieces: 1}, {ok: true, pieces: 2, affixes: 2}} {
		if p, _ := m.collectPieces(ctx, rs, append(st[:3:3], end)); p != nil {
			t.Errorf("%+v: %v", end, p)
		}
	}
	short := []segState{{ok: true}, {ok: true, pieces: 1, known: true}, {ok: true, pieces: 2, prev: 1, known: true}}
	if p, _ := m.collectPieces(ctx, []rune("땅물"), short); p != nil {
		t.Fatal("one-syllable pieces only", p)
	}
	if _, err := newTestMemo(t, "류").collectPieces(ctx, []rune("오류"), short); err == nil {
		t.Fatal("want error")
	}
}
