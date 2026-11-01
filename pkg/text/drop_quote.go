//ff:func feature=text type=helper control=sequence
//ff:what dropQuote — strings.Map용: 따옴표면 지운다(−1)
package text

import "strings"

// dropQuote removes quotation marks when used with strings.Map.
func dropQuote(r rune) rune {
	if strings.ContainsRune(quoteRunes, r) {
		return -1
	}
	return r
}
