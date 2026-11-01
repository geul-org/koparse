//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what bestEnding — i에서 끝나는 조각을 모두 따져 가장 나은 칸(어간 전체는 조각이 아님, 접사는 가장자리 한 음절만)
package noun

import "context"

// bestEnding returns the best state for the prefix rs[:i].
func (m *Memo) bestEnding(ctx context.Context, rs []rune, st []segState, i int, nominal bool) (segState, error) {
	best := segState{}
	for j := 0; j < i; j++ {
		if !st[j].ok || (j == 0 && i == len(rs)) {
			continue
		}
		known, err := m.knownPiece(ctx, string(rs[j:i]))
		if err != nil {
			return best, err
		}
		s := nextState(st[j], j, known, nominal && i-j == 1 && (j == 0 || i == len(rs)))
		if betterState(s, best) {
			best = s
		}
	}
	return best, nil
}
