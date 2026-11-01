//ff:func feature=stats type=helper control=sequence
//ff:what josaAfter — 명사 토큰 바로 뒤 같은 어절의 조사(없으면 ∅)
package stats

// josaAfter returns the particle token right after toks[p] in the same eojeol, or "∅".
func josaAfter(toks []Token, p int) string {
	if p+1 < len(toks) && toks[p+1].Tag == "J" && toks[p+1].Eojeol == toks[p].Eojeol {
		return toks[p+1].Text
	}
	return "∅"
}
