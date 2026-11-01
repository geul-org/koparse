//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what collectPieces — 동적 계획 끝 칸에서 거슬러 조각을 모은다(조각 둘 이상, 두 음절 이상 아는 조각이 있을 때만)
package noun

import "context"

// collectPieces walks back from the last state and returns the pieces, or nil
// when the stem does not split into at least two pieces with a known piece
// of two or more syllables (발달 is not 발+달).
func (m *Memo) collectPieces(ctx context.Context, rs []rune, st []segState) ([]piece, error) {
	end := st[len(rs)]
	if !end.ok || end.pieces < 2 || end.affixes == end.pieces {
		return nil, nil
	}
	out := make([]piece, end.pieces)
	for i, k := len(rs), end.pieces-1; i > 0; k-- {
		s := st[i]
		r, err := m.lookup(ctx, string(rs[s.prev:i]))
		if err != nil {
			return nil, err
		}
		out[k] = piece{form: string(rs[s.prev:i]), known: s.known, r: r}
		i = s.prev
	}
	if !hasLongKnown(out) {
		return nil, nil
	}
	return out, nil
}
