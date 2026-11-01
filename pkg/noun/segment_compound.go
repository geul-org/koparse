//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what segmentCompound — 사전에 없는 한글 어간을 아는 조각들로 나눈다(조각 최소, 접미가 붙은 어절만 가장자리 한 음절 접사 허용)
package noun

import (
	"context"

	"github.com/geul-org/koparse/pkg/hangul"
)

// segmentCompound splits an unknown Hangul stem into known pieces (반사신경 →
// 반사+신경). Edge one-syllable affixes (전-, -성) are allowed only when the
// eojeol carries a suffix (nominal). Returns nil when no segmentation exists.
func (m *Memo) segmentCompound(ctx context.Context, stem string, nominal bool) ([]piece, error) {
	rs := []rune(stem)
	if !hangul.IsAllHangul(stem) || len(rs) < 2 || len(rs) > maxCompound {
		return nil, nil
	}
	st := make([]segState, len(rs)+1)
	st[0].ok = true
	for i := 1; i <= len(rs); i++ {
		s, err := m.bestEnding(ctx, rs, st, i, nominal)
		if err != nil {
			return nil, err
		}
		st[i] = s
	}
	return m.collectPieces(ctx, rs, st)
}
