//ff:func feature=text type=helper control=sequence
//ff:what appendTrimmed — 앞뒤 공백을 떼고 비어 있지 않으면 붙인다
package text

import "strings"

// appendTrimmed appends strings.TrimSpace(s) to out unless it is empty.
func appendTrimmed(out []string, s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return out
	}
	return append(out, s)
}
