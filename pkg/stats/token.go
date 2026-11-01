//ff:type feature=stats type=model
//ff:what Token — 조사를 뗀 토큰(글, 품사 N·J·V·M, 어절 번호, 절 번호, 명사 단위 번호 또는 −1)
package stats

// Token is one token of the particle-detached stream used for distances.
type Token struct {
	Text   string
	Tag    string // N noun, J particles, V predicate (incl. copula/verbal suffix), M other
	Eojeol int
	Clause int
	Unit   int // index into the sentence's noun units, −1 if not a noun
}
