//ff:func feature=analyze type=builder control=sequence
//ff:what sentenceUnits — 문장 하나를 분리하고 명사 단위를 찾는다
package analyze

import (
	"context"

	"github.com/geul-org/koparse/pkg/text"
)

// sentenceUnits splits s and finds its noun units.
func (a *Analyzer) sentenceUnits(ctx context.Context, s text.Sentence) (sentenceWork, error) {
	w, err := a.splitSentence(ctx, s)
	if err != nil {
		return w, err
	}
	w.units, err = a.memo.FindUnits(ctx, w.words)
	return w, err
}
