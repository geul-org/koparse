//ff:func feature=stats type=builder control=iteration dimension=1
//ff:what buildTokens — 문장을 조사를 뗀 토큰열로: 명사 단위 N, 그 밖의 어간 V·M, 접미 J·V
package stats

import "github.com/geul-org/koparse/pkg/noun"

// buildTokens builds the token stream of one sentence.
func buildTokens(words []noun.Word, units []noun.Unit) []Token {
	cov := coverage(len(words), units)
	var toks []Token
	clause := 0
	for e, w := range words {
		toks = unitTokens(toks, units, e, clause)
		if !cov[e] {
			toks = append(toks, stemToken(w, e, clause))
		}
		toks = suffixTokens(toks, w, e, clause)
		clause = nextClause(clause, w)
	}
	return toks
}
