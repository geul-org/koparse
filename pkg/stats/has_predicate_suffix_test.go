//ff:func feature=stats type=helper control=sequence
//ff:what TestHasPredicateSuffix — 서술격·용언화 접미가 있으면 참, 조사뿐이면 거짓
package stats

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestHasPredicateSuffix(t *testing.T) {
	p := lexicon.Suffix{Form: "에", Kind: lexicon.KindParticle}
	c := lexicon.Suffix{Form: "이다", Kind: lexicon.KindCopula}
	if !hasPredicateSuffix(noun.Word{Chain: []lexicon.Suffix{p, c}}) || hasPredicateSuffix(noun.Word{Chain: []lexicon.Suffix{p}}) {
		t.Fatal("hasPredicateSuffix")
	}
}
