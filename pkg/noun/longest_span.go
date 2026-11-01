//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what longestSpan — i에서 시작하는 여러 어절 구간을 긴 것부터 띄어 쓴 형태·붙여 쓴 형태로 찾는다
package noun

import (
	"context"
	"strings"
)

// longestSpan looks up spans words[i..j] (j > i, at most maxSpan eojeols)
// longest first, as written with spaces and joined.
func (m *Memo) longestSpan(ctx context.Context, words []Word, i, b int) (Unit, bool, error) {
	for j := min(b, i+maxSpan-1); j > i; j-- {
		spaced := spanText(words, i, j, " ")
		form, r, err := m.firstFound(ctx, []string{spaced, strings.ReplaceAll(spaced, " ", "")})
		if err != nil {
			return Unit{}, false, err
		}
		if form != "" {
			return makeUnit(words, i, j, spaced, form, r, false, true), true, nil
		}
	}
	return Unit{}, false, nil
}
