//ff:func feature=noun type=judge control=sequence
//ff:what JudgeDoc — 문서의 명사 단위를 두 번 판정한다: 1차 채택으로 이웃 분야를 모으고 2차에 R4b로 쓴다
package noun

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// JudgeDoc judges all units of one document. The first pass (no neighbors)
// adopts by R1–R4a; its adoptions become the neighbor field of the second.
func JudgeDoc(units []Unit, dc DocContext, lex *lexicon.Lexicon, cfg config.Noun) ([]Result, error) {
	dc.Neighbors = nil
	first, err := judgeAll(units, dc, lex, cfg)
	if err != nil {
		return nil, err
	}
	dc.Neighbors = neighborsOf(units, first)
	return judgeAll(units, dc, lex, cfg)
}
