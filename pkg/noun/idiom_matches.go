//ff:func feature=noun type=rule control=iteration dimension=1
//ff:what idiomMatches — 명사 단위가 관용 틀 하나에라도 드는지(명사, 조사, 지배 서술어 앞부분)
package noun

import "github.com/geul-org/koparse/pkg/lexicon"

// idiomMatches reports whether u sits in any idiom frame.
func idiomMatches(idioms []lexicon.Idiom, u Unit) bool {
	for _, id := range idioms {
		if (id.Noun == "*" || id.Noun == u.Form) && particleFits(id.Particles, u.Chain) && hasAnyPrefix(u.Governor, id.Governors) {
			return true
		}
	}
	return false
}
