//ff:func feature=stats type=builder control=sequence
//ff:what appendJ — 조사 글이 있으면 J 토큰을 붙인다
package stats

// appendJ appends a J token for j unless it is empty.
func appendJ(toks []Token, j string, e, clause int) []Token {
	if j == "" {
		return toks
	}
	return append(toks, Token{Text: j, Tag: "J", Eojeol: e, Clause: clause, Unit: -1})
}
