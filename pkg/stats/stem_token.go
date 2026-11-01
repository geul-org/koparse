//ff:func feature=stats type=builder control=sequence
//ff:what stemToken — 명사가 아닌 어절의 어간 토큰(서술어로 보이면 V, 아니면 M)
package stats

import "github.com/geul-org/koparse/pkg/noun"

// stemToken returns the token of a non-noun stem.
func stemToken(w noun.Word, e, clause int) Token {
	tag := "M"
	if len(w.Chain) == 0 && isPredicate(w.Surface) {
		tag = "V"
	}
	return Token{Text: w.Stem, Tag: tag, Eojeol: e, Clause: clause, Unit: -1}
}
