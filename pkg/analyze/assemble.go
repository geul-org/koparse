//ff:func feature=analyze type=builder control=iteration dimension=1
//ff:what assemble — 문장 하나의 출력 줄을 만든다(어절 분리, 명사 판정, 문맥 통계)
package analyze

import (
	"github.com/geul-org/koparse/pkg/noun"
	"github.com/geul-org/koparse/pkg/stats"
)

// assemble builds the output line of one sentence; nouns[i] judges w.units[i].
func (a *Analyzer) assemble(w sentenceWork, nouns []noun.Result) SentenceOut {
	cov := coveredEojeols(w.units)
	out := SentenceOut{Text: w.sent.Text, Asides: w.sent.Asides, Nouns: nouns, Stats: stats.Records(w.words, w.units, nouns)}
	if out.Nouns == nil {
		out.Nouns = []noun.Result{}
	}
	for i, d := range w.decisions {
		e := EojeolOut{Surface: w.words[i].Surface, Split: eojeolParts(w.words[i], cov[i]), Verdict: d.Ranked[0].Verdict, Status: d.Status}
		if a.Trace {
			e.Ranked = d.Ranked[:min(len(d.Ranked), 5)]
		}
		out.Eojeol = append(out.Eojeol, e)
	}
	return out
}
