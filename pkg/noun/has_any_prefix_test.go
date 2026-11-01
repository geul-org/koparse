//ff:func feature=noun type=helper control=sequence
//ff:what TestHasAnyPrefix — 앞부분 하나라도 맞으면 참, 빈 문자열은 거짓
package noun

import "testing"

func TestHasAnyPrefix(t *testing.T) {
	if !hasAnyPrefix("삼아", []string{"하", "삼"}) || hasAnyPrefix("먹고", []string{"하"}) || hasAnyPrefix("", []string{""}) {
		t.Fatal("hasAnyPrefix")
	}
}
