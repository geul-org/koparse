//ff:func feature=text type=helper control=sequence
//ff:what headBefore — 괄호 앞 낱말(문장부호 뗌)
package text

import "strings"

// headBefore returns the last word of b without surrounding punctuation.
func headBefore(b []rune) string {
	f := strings.Fields(string(b))
	if len(f) == 0 {
		return ""
	}
	return strings.TrimFunc(f[len(f)-1], isTrimRune)
}
