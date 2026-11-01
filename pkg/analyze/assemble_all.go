//ff:func feature=analyze type=builder control=iteration dimension=1
//ff:what assembleAll — 문서 전체 판정 결과를 문장별로 나눠 출력 줄을 만든다
package analyze

import "github.com/geul-org/koparse/pkg/noun"

// assembleAll splits the flat results back per sentence.
func (a *Analyzer) assembleAll(works []sentenceWork, results []noun.Result) []SentenceOut {
	out := make([]SentenceOut, 0, len(works))
	off := 0
	for _, w := range works {
		out = append(out, a.assemble(w, results[off:off+len(w.units)]))
		off += len(w.units)
	}
	return out
}
