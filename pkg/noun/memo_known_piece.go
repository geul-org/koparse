//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what knownPiece — 복합어 조각으로 인정할지: 두 음절 이상은 후보가 있으면, 한 음절은 고유 후보(음절·이름 등)가 아닌 후보가 있어야
package noun

import (
	"context"

	"github.com/geul-org/koparse/pkg/hangul"
)

// knownPiece reports whether s may be a piece of a compound. A one-syllable
// piece needs a candidate outside the proper classes (Hangul syllable and
// name items would make every syllable "known").
func (m *Memo) knownPiece(ctx context.Context, s string) (bool, error) {
	r, err := m.lookup(ctx, s)
	if err != nil || hangul.Count(s) > 1 {
		return len(r.cands) > 0, err
	}
	for _, e := range r.infos {
		if !anyIn(e.P31, m.Lex.Proper) {
			return true, nil
		}
	}
	return false, nil
}
