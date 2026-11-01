//ff:func feature=noun type=builder control=sequence
//ff:what singleUnits — 어절 하나: 어간(공백 변형 포함) → 복합어 조각 → 접미가 붙었으면 후보 없는 명사 → 아니면 명사 아님
package noun

import (
	"context"

	"github.com/geul-org/koparse/pkg/hangul"
)

// singleUnits finds the units of one eojeol's stem.
func (m *Memo) singleUnits(ctx context.Context, words []Word, i int) ([]Unit, error) {
	w := words[i]
	form, r, err := m.firstFound(ctx, spaceVariants(w.Stem))
	if err != nil {
		return nil, err
	}
	if form != "" {
		return []Unit{makeUnit(words, i, i, w.Stem, form, r, false, true)}, nil
	}
	pieces, err := m.segmentCompound(ctx, w.Stem, len(w.Chain) > 0)
	if err != nil {
		return nil, err
	}
	if pieces != nil {
		return pieceUnits(words, i, pieces), nil
	}
	if len(w.Chain) > 0 && hangul.IsAllHangul(w.Stem) {
		return []Unit{makeUnit(words, i, i, w.Stem, "", lookupResult{}, false, true)}, nil
	}
	return nil, nil
}
