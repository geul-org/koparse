//ff:func feature=text type=helper control=sequence
//ff:what TestDropQuote — 따옴표는 지우고 글자는 둔다
package text

import "testing"

func TestDropQuote(t *testing.T) {
	if dropQuote('"') != -1 || dropQuote('“') != -1 || dropQuote('가') != '가' {
		t.Fatal("dropQuote wrong")
	}
}
