//ff:func feature=text type=helper control=sequence
//ff:what isMiddleDot — 가운뎃점(·, ㆍ, ・)인지. 어절 안 나열을 가른다
package text

// isMiddleDot reports whether r is a middle dot used to list words (경기도·인천광역시).
func isMiddleDot(r rune) bool {
	return r == '·' || r == 'ㆍ' || r == '・'
}
