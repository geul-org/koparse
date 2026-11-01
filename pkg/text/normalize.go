//ff:func feature=text type=helper control=sequence
//ff:what Normalize — 유니코드 NFC로 맞춘다
package text

import "golang.org/x/text/unicode/norm"

// Normalize returns s in Unicode NFC.
func Normalize(s string) string {
	return norm.NFC.String(s)
}
