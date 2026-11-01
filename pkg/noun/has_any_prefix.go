//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what hasAnyPrefix — 문자열이 앞부분 목록 중 하나로 시작하는지(빈 문자열은 아니다)
package noun

import "strings"

// hasAnyPrefix reports whether non-empty s starts with any of prefixes.
func hasAnyPrefix(s string, prefixes []string) bool {
	for _, p := range prefixes {
		if s != "" && strings.HasPrefix(s, p) {
			return true
		}
	}
	return false
}
