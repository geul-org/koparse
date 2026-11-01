//ff:func feature=text type=parser control=iteration dimension=1
//ff:what Tokenize — 공백과 가운뎃점으로 어절을 나눈다
package text

import "strings"

// Tokenize splits a sentence body into eojeols at whitespace and middle dots.
func Tokenize(s string) []Eojeol {
	var out []Eojeol
	for _, field := range strings.Fields(s) {
		out = attachField(out, field)
	}
	return out
}
