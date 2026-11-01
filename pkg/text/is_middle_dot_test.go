//ff:func feature=text type=helper control=sequence
//ff:what TestIsMiddleDot — 가운뎃점 세 가지만
package text

import "testing"

func TestIsMiddleDot(t *testing.T) {
	if !isMiddleDot('·') || !isMiddleDot('ㆍ') || !isMiddleDot('・') || isMiddleDot('.') {
		t.Fatal("isMiddleDot wrong")
	}
}
