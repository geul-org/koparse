//ff:func feature=eval type=builder control=iteration dimension=1
//ff:what RunSamples — 표본마다 주제 QID로 분석해 문장 열쇠(item.sent, 1부터)별 출력을 모은다
package eval

import (
	"context"
	"fmt"

	"github.com/geul-org/koparse/pkg/analyze"
)

// RunSamples analyzes every sample with its topic and returns the outputs
// keyed "item.sent" (sentences numbered from 1).
func RunSamples(ctx context.Context, a *analyze.Analyzer, samples []Sample) (map[string]analyze.SentenceOut, error) {
	outs := map[string]analyze.SentenceOut{}
	for _, s := range samples {
		sents, err := a.Document(ctx, s.Text, s.QID)
		if err != nil {
			return nil, fmt.Errorf("sample %d: %w", s.ID, err)
		}
		addSentences(outs, s.ID, sents)
	}
	return outs, nil
}
