//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what pieceUnits — 아는 조각마다 단위를 만든다(마지막 아는 조각이 어절의 접미·지배어를 갖는다)
package noun

// pieceUnits turns the known pieces of words[i]'s stem into units.
func pieceUnits(words []Word, i int, pieces []piece) []Unit {
	last := -1
	for k, p := range pieces {
		if p.known {
			last = k
		}
	}
	var out []Unit
	for k, p := range pieces {
		if p.known {
			out = append(out, makeUnit(words, i, i, p.form, p.form, p.r, true, k == last))
		}
	}
	return out
}
