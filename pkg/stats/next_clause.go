//ff:func feature=stats type=helper control=sequence
//ff:what nextClause — 어절 뒤 절 번호: 쉼표류 또는 연결어미로 끝나는 서술어 뒤에서 하나 늘린다
package stats

import "github.com/geul-org/koparse/pkg/noun"

// nextClause returns the clause number after w.
func nextClause(clause int, w noun.Word) int {
	pred := hasPredicateSuffix(w) || (len(w.Chain) == 0 && isPredicate(w.Surface))
	if w.Break || (pred && connectiveSyllables[lastSyllable(w.Surface)]) {
		return clause + 1
	}
	return clause
}
