//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what FindUnits — 문장의 명사 단위를 찾는다: 맨 어절 묶음마다 긴 구간부터, 없으면 어간·공백 변형·복합어 조각
package noun

import "context"

// FindUnits finds the noun units of one sentence. Consecutive bare eojeols
// (no suffix, no comma) and the eojeol ending them form a run; inside a run
// multi-eojeol spans are tried longest first (복합어는 통째로 먼저 찾고).
func (m *Memo) FindUnits(ctx context.Context, words []Word) ([]Unit, error) {
	var out []Unit
	for a := 0; a < len(words); {
		b := runEnd(words, a)
		us, err := m.runUnits(ctx, words, a, b)
		if err != nil {
			return nil, err
		}
		out = append(out, us...)
		a = b + 1
	}
	return out, nil
}
