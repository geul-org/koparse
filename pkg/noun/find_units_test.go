//ff:func feature=noun type=builder control=sequence
//ff:what TestFindUnits — 지은 문장: 여러 어절 구간, 공백 변형, 복합어 조각, 후보 없는 명사, 명사 아닌 어절, 오류
package noun

import (
	"context"
	"reflect"
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestFindUnits(t *testing.T) {
	m := newTestMemo(t, "오류")
	p := []lexicon.Suffix{{Form: "은", Kind: "particle"}}
	words := []Word{
		{Surface: "강", Stem: "강"}, {Surface: "하구는", Stem: "하구", Chain: p},
		{Surface: "발효식품은", Stem: "발효식품", Chain: p},
		{Surface: "김치땅은", Stem: "김치땅", Chain: p},
		{Surface: "모름은", Stem: "모름", Chain: p, Break: true},
		{Surface: "달린다", Stem: "달린다"},
	}
	us, err := m.FindUnits(context.Background(), words)
	want := []string{"강 하구", "발효식품", "김치", "땅", "모름"}
	if err != nil || !reflect.DeepEqual(unitForms(us), want) {
		t.Fatal(unitForms(us), err)
	}
	if _, err := m.FindUnits(context.Background(), []Word{{Surface: "오류", Stem: "오류"}}); err == nil {
		t.Fatal("want error")
	}
}
