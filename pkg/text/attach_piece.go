//ff:func feature=text type=parser control=sequence
//ff:what attachPiece — 조각 하나를 어절로 붙인다(문장부호만 남은 조각은 앞 어절에 쉼표류만 옮긴다)
package text

// attachPiece appends piece as an eojeol, or, when it is punctuation only,
// carries its clause break over to the previous eojeol.
func attachPiece(out []Eojeol, piece string) []Eojeol {
	e, ok := makeEojeol(piece)
	if ok {
		return append(out, e)
	}
	if e.Break && len(out) > 0 {
		out[len(out)-1].Break = true
	}
	return out
}
