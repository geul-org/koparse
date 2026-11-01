//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what TestSingleUnits — 어간, 공백 변형(R5), 조각, 접미 붙은 모르는 명사, 명사 아님, 비한글 어간, 오류
package noun

import (
	"context"
	"reflect"
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestSingleUnits(t *testing.T) {
	m := newTestMemo(t, "오류")
	p := []lexicon.Suffix{{Form: "을", Kind: "particle"}}
	cases := []struct {
		w     Word
		forms []string
	}{
		{Word{Stem: "김치", Chain: p}, []string{"김치"}},
		{Word{Stem: "발효식품", Chain: p}, []string{"발효식품"}},
		{Word{Stem: "김치땅", Chain: p}, []string{"김치", "땅"}},
		{Word{Stem: "모름", Chain: p}, []string{"모름"}},
		{Word{Stem: "달린다"}, nil},
		{Word{Stem: "km2", Chain: p}, nil},
	}
	for _, c := range cases {
		us, err := m.singleUnits(context.Background(), []Word{c.w}, 0)
		if err != nil || !reflect.DeepEqual(unitForms(us), c.forms) {
			t.Errorf("%s: %v %v", c.w.Stem, unitForms(us), err)
		}
	}
	if _, err := m.singleUnits(context.Background(), []Word{{Stem: "오류"}}, 0); err == nil {
		t.Fatal("want lookup error")
	}
	m2 := newTestMemo(t, "김")
	if _, err := m2.singleUnits(context.Background(), []Word{{Stem: "김치땅"}}, 0); err == nil {
		t.Fatal("want segmentation error")
	}
}
