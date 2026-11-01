//ff:func feature=stats type=builder control=iteration dimension=1
//ff:what pairsFor — 명사 토큰과 문장 안 다른 명사·조사·서술어의 거리 쌍
package stats

// pairsFor returns the pairs of the noun at p with every other N, J, V token.
func pairsFor(toks []Token, p int) []Pair {
	out := []Pair{}
	for q, t := range toks {
		if q == p || t.Tag == "M" {
			continue
		}
		out = append(out, Pair{Other: t.Text, Tag: t.Tag, Bucket: Bucket(q - p),
			SameEojeol: t.Eojeol == toks[p].Eojeol, CrossClause: t.Clause != toks[p].Clause})
	}
	return out
}
