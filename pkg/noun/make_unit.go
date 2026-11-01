//ff:func feature=noun type=builder control=sequence
//ff:what makeUnit — 명사 단위를 만든다(어간 끝에 닿는 단위만 어절의 접미·지배어를 갖는다)
package noun

// makeUnit builds a unit over words[first..last]. Only a final unit (one
// that ends where the stem ends) carries the eojeol's suffixes and governor.
func makeUnit(words []Word, first, last int, form, matched string, r lookupResult, piece, final bool) Unit {
	u := Unit{Form: form, Matched: matched, First: first, Last: last, Piece: piece, Final: final, Cands: r.cands, Infos: r.infos}
	if final {
		u.Chain = chainForms(words[last].Chain)
		u.Governor = governorOf(words, last)
	}
	return u
}
