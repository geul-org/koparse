//ff:func feature=noun type=rule control=sequence
//ff:what TestIdiomMatches — 아무 명사·특정 명사 틀, 조사나 지배어가 다르면 거짓
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestIdiomMatches(t *testing.T) {
	ids := []lexicon.Idiom{{Noun: "*", Particles: []string{"으로", "로"}, Governors: []string{"삼"}}, {Noun: "자리", Particles: []string{"-"}, Governors: []string{"잡"}}}
	if !idiomMatches(ids, Unit{Form: "원칙", Chain: []string{"으로"}, Governor: "삼아"}) ||
		!idiomMatches(ids, Unit{Form: "자리", Governor: "잡고"}) ||
		idiomMatches(ids, Unit{Form: "자리", Chain: []string{"를"}, Governor: "잡고"}) ||
		idiomMatches(ids, Unit{Form: "원칙", Chain: []string{"으로"}, Governor: "먹어"}) {
		t.Fatal("idiomMatches")
	}
}
