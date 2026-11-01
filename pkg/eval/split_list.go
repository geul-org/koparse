//ff:func feature=eval type=helper control=sequence
//ff:what splitList — 쉼표 목록을 나눈다(빈 문자열은 빈 목록)
package eval

import "strings"

// splitList splits a comma list; "" gives nil.
func splitList(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	return strings.Split(s, ",")
}
