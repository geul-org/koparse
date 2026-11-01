//ff:func feature=analyze type=builder control=iteration dimension=1
//ff:what Document — 문서 하나: 정규화·문장 → 어절 분리 → 명사 단위 → 문서 단위 명사 판정(2회) → 문장별 출력·통계
package analyze

import (
	"context"

	"github.com/geul-org/koparse/pkg/noun"
	"github.com/geul-org/koparse/pkg/text"
)

// Document analyzes one document. topic is the QID of the document's subject
// ("" if unknown); it feeds R2 and R3.
func (a *Analyzer) Document(ctx context.Context, input, topic string) ([]SentenceOut, error) {
	stmts, err := a.statementSet(ctx, topic)
	if err != nil {
		return nil, err
	}
	var works []sentenceWork
	var units []noun.Unit
	for _, s := range text.Parse(input) {
		w, err := a.sentenceUnits(ctx, s)
		if err != nil {
			return nil, err
		}
		works = append(works, w)
		units = append(units, w.units...)
	}
	results, err := noun.JudgeDoc(units, noun.DocContext{Topic: topic, Statements: stmts}, a.Lex, a.Cfg.Noun)
	if err != nil {
		return nil, err
	}
	return a.assembleAll(works, results), nil
}
