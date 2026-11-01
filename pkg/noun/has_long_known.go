//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what hasLongKnown — 두 음절 이상인 아는 조각이 있는지(한 음절 조각만으로는 복합어로 보지 않는다)
package noun

import "github.com/geul-org/koparse/pkg/hangul"

// hasLongKnown reports whether some known piece has two or more syllables.
// A stem made only of one-syllable pieces is more likely an unlisted word
// (발달) than a compound (발+달).
func hasLongKnown(pieces []piece) bool {
	for _, p := range pieces {
		if p.known && hangul.Count(p.form) > 1 {
			return true
		}
	}
	return false
}
