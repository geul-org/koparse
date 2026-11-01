//ff:func feature=text type=parser control=iteration dimension=1
//ff:what attachField — 공백으로 나뉜 덩어리를 가운뎃점으로 다시 나눠 어절로 붙인다
package text

import "strings"

// attachField splits one whitespace field at middle dots and appends the pieces.
func attachField(out []Eojeol, field string) []Eojeol {
	for _, piece := range strings.FieldsFunc(field, isMiddleDot) {
		out = attachPiece(out, piece)
	}
	return out
}
