//ff:func feature=hangul type=helper control=sequence
//ff:what TestIsSyllable — 가·힣은 음절, ㄱ·A는 아니다
package hangul

import "testing"

func TestIsSyllable(t *testing.T) {
	if !IsSyllable('가') || !IsSyllable('힣') || IsSyllable('ㄱ') || IsSyllable('A') {
		t.Fatal("IsSyllable boundaries wrong")
	}
}
