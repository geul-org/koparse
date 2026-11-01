//ff:func feature=split type=judge control=iteration dimension=1
//ff:what Split — 어절 하나: 후보 생성 → 어간마다 지식 후보 수 → TAN 판정
package split

import (
	"context"

	"github.com/geul-org/koparse/pkg/text"
)

// Split judges the split of one eojeol.
func (s *Splitter) Split(ctx context.Context, e text.Eojeol) (Decision, error) {
	cands := Generate(s.Lex, e.Surface)
	counts := make([]int, len(cands))
	for i, c := range cands {
		n, err := s.stemCount(ctx, c.Stem)
		if err != nil {
			return Decision{}, err
		}
		counts[i] = n
	}
	return judge(&facts{cands: cands, stemCount: counts, quoted: e.Quoted, lex: s.Lex, cfg: s.Cfg})
}
