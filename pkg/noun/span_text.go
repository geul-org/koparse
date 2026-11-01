//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what spanText — 어절 i..j의 어간을 sep로 잇는다(앞 어절들은 맨 어절이라 어간 = 표면)
package noun

import "strings"

// spanText joins the stems of words[i..j] with sep.
func spanText(words []Word, i, j int, sep string) string {
	parts := make([]string, 0, j-i+1)
	for k := i; k <= j; k++ {
		parts = append(parts, words[k].Stem)
	}
	return strings.Join(parts, sep)
}
