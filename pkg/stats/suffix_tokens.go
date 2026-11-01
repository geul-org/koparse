//ff:func feature=stats type=builder control=iteration dimension=1
//ff:what suffixTokens — 접미 사슬을 토큰으로: 이어진 조사는 J 하나, 서술격·용언화 접미는 V
package stats

import (
	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

// suffixTokens appends w's suffixes: consecutive particles as one J token,
// a copula or verbal suffix as a V token.
func suffixTokens(toks []Token, w noun.Word, e, clause int) []Token {
	j := ""
	for _, s := range w.Chain {
		if s.Kind == lexicon.KindParticle {
			j += s.Form
			continue
		}
		toks = appendJ(toks, j, e, clause)
		j = ""
		toks = append(toks, Token{Text: s.Form, Tag: "V", Eojeol: e, Clause: clause, Unit: -1})
	}
	return appendJ(toks, j, e, clause)
}
