//ff:func feature=analyze type=builder control=iteration dimension=1
//ff:what splitSentence — 문장의 어절마다 분리를 판정하고 1위 분리로 어절을 만든다
package analyze

import (
	"context"

	"github.com/geul-org/koparse/pkg/noun"
	"github.com/geul-org/koparse/pkg/text"
)

// splitSentence splits every eojeol of s.
func (a *Analyzer) splitSentence(ctx context.Context, s text.Sentence) (sentenceWork, error) {
	w := sentenceWork{sent: s}
	for _, e := range s.Eojeols {
		d, err := a.splitter.Split(ctx, e)
		if err != nil {
			return w, err
		}
		best := d.Ranked[0].Candidate
		w.decisions = append(w.decisions, d)
		w.words = append(w.words, noun.Word{Surface: e.Surface, Stem: best.Stem, Chain: best.Suffixes, Break: e.Break})
	}
	return w, nil
}
