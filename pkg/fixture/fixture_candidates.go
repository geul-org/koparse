//ff:func feature=fixture type=adapter control=iteration dimension=1
//ff:what Candidates — 표면이 라벨·별칭과 정확히 같은 후보(라벨 먼저). 표에 없는 표면은 놓침으로 센다
package fixture

import (
	"context"

	"github.com/geul-org/koparse/pkg/knowledge"
)

// Candidates returns label matches then alias matches (an entity matching
// both is reported once, as a label match).
func (f *Fixture) Candidates(_ context.Context, surface string, lang string) ([]knowledge.Candidate, error) {
	e, ok := f.forms[lang][surface]
	if !ok {
		f.mu.Lock()
		f.misses[lang+":"+surface]++
		f.mu.Unlock()
		return nil, nil
	}
	seen := map[string]bool{}
	var out []knowledge.Candidate
	for i, q := range append(append([]string{}, e.Label...), e.Alias...) {
		if seen[q] {
			continue
		}
		seen[q] = true
		m := knowledge.MatchLabel
		if i >= len(e.Label) {
			m = knowledge.MatchAlias
		}
		out = append(out, knowledge.Candidate{QID: q, Match: m, Surface: surface, Lang: lang})
	}
	return out, nil
}
