//ff:func feature=text type=helper control=sequence
//ff:what TestIsTrimRune — 쉼표·따옴표는 떼고 %·-·글자는 남긴다
package text

import "testing"

func TestIsTrimRune(t *testing.T) {
	if !isTrimRune(',') || !isTrimRune('"') || !isTrimRune('“') || isTrimRune('%') || isTrimRune('-') || isTrimRune('가') {
		t.Fatal("isTrimRune wrong")
	}
}
