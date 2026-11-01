//ff:func feature=stats type=helper control=iteration dimension=1
//ff:what hasPredicateSuffix — 접미 사슬에 서술격·용언화 접미가 있는지
package stats

import (
	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

// hasPredicateSuffix reports whether w's chain contains a copula or verbal suffix.
func hasPredicateSuffix(w noun.Word) bool {
	for _, s := range w.Chain {
		if s.Kind != lexicon.KindParticle {
			return true
		}
	}
	return false
}
