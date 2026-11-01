//ff:func feature=stats type=helper control=iteration dimension=1
//ff:what negatedAt — 어절 e(지배 서술어)가 부정인지: 어절 안 또는 다음 어절 머리에 않·못·아니·없
package stats

import (
	"strings"

	"github.com/geul-org/koparse/pkg/noun"
)

// negatedAt reports whether the predicate eojeol e is negated.
func negatedAt(words []noun.Word, e int) bool {
	for _, m := range negationMarks {
		if strings.Contains(words[e].Surface, m) || (e+1 < len(words) && strings.HasPrefix(words[e+1].Surface, m)) {
			return true
		}
	}
	return false
}
