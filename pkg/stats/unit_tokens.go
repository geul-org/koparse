//ff:func feature=stats type=builder control=iteration dimension=1
//ff:what unitTokens — 어절 e에서 시작하는 명사 단위마다 N 토큰(여러 어절 단위는 마지막 어절에 둔다)
package stats

import "github.com/geul-org/koparse/pkg/noun"

// unitTokens appends an N token for every unit starting at eojeol e.
func unitTokens(toks []Token, units []noun.Unit, e, clause int) []Token {
	for i, u := range units {
		if u.First == e {
			toks = append(toks, Token{Text: u.Form, Tag: "N", Eojeol: u.Last, Clause: clause, Unit: i})
		}
	}
	return toks
}
