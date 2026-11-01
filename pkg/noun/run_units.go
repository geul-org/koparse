//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what runUnits — 묶음 안에서 왼쪽부터 가장 긴 구간 단위, 없으면 어절 하나 단위
package noun

import "context"

// runUnits finds units in words[a..b].
func (m *Memo) runUnits(ctx context.Context, words []Word, a, b int) ([]Unit, error) {
	var out []Unit
	for i := a; i <= b; {
		u, ok, err := m.longestSpan(ctx, words, i, b)
		if err != nil {
			return nil, err
		}
		if ok {
			out = append(out, u)
			i = u.Last + 1
			continue
		}
		us, err := m.singleUnits(ctx, words, i)
		if err != nil {
			return nil, err
		}
		out = append(out, us...)
		i++
	}
	return out, nil
}
