//ff:func feature=stats type=helper control=sequence
//ff:what TestNextClause — 쉼표, 연결어미 서술어(통째·서술격)에서 늘고, 조사 어절·종결 서술어에서는 그대로
package stats

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestNextClause(t *testing.T) {
	ija := []lexicon.Suffix{{Form: "이자", Kind: lexicon.KindCopula}}
	go_ := []lexicon.Suffix{{Form: "고", Kind: lexicon.KindParticle}}
	if nextClause(0, noun.Word{Surface: "가", Break: true}) != 1 ||
		nextClause(0, noun.Word{Surface: "먹고"}) != 1 ||
		nextClause(0, noun.Word{Surface: "수도이자", Chain: ija}) != 1 ||
		nextClause(0, noun.Word{Surface: "사과고", Chain: go_}) != 0 ||
		nextClause(0, noun.Word{Surface: "먹는다"}) != 0 {
		t.Fatal("nextClause")
	}
}
