//ff:func feature=eval type=helper control=sequence
//ff:what itemOf — 문장 열쇠 item.sent의 표본 번호 부분
package eval

import "strings"

// itemOf returns the item part of "item.sent".
func itemOf(key string) string {
	item, _, _ := strings.Cut(key, ".")
	return item
}
