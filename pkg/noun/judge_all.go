//ff:func feature=noun type=judge control=iteration dimension=1
//ff:what judgeAll — 명사 단위 모두를 같은 문서 문맥으로 판정한다
package noun

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// judgeAll judges every unit with the same document context.
func judgeAll(units []Unit, dc DocContext, lex *lexicon.Lexicon, cfg config.Noun) ([]Result, error) {
	out := make([]Result, len(units))
	for i, u := range units {
		r, err := judgeUnit(u, dc, lex, cfg)
		if err != nil {
			return nil, err
		}
		out[i] = r
	}
	return out, nil
}
