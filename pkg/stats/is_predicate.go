//ff:func feature=stats type=helper control=selection
//ff:what isPredicate — 명사가 아닌 통째 어절이 서술어로 보이는지(끝 글자가 어미 글자이거나 관형형 ㄴ·ㄹ 받침)
package stats

import "github.com/geul-org/koparse/pkg/hangul"

// isPredicate guesses whether a non-noun eojeol is a predicate from its last
// syllable: a common ending syllable, or a final ㄴ/ㄹ (떨어진, 배울).
func isPredicate(s string) bool {
	r := lastSyllable(s)
	switch {
	case !hangul.IsSyllable(r):
		return false
	case endingSyllables[r]:
		return true
	default:
		f := (r - 0xAC00) % 28
		return f == 4 || f == 8
	}
}
