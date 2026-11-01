//ff:func feature=split type=helper control=sequence
//ff:what TestHasHangul — 한글 음절이 하나라도 있으면 참
package split

import "testing"

func TestHasHangul(t *testing.T) {
	if !hasHangul("5,066명") || hasHangul("km2") || hasHangul("") {
		t.Fatal("hasHangul")
	}
}
