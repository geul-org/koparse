//ff:func feature=text type=helper control=sequence
//ff:what TestDropTrailingSpace — 끝 공백 하나만 떼고 빈 입력은 그대로
package text

import "testing"

func TestDropTrailingSpace(t *testing.T) {
	if string(dropTrailingSpace([]rune("가  "))) != "가 " || string(dropTrailingSpace([]rune("가"))) != "가" || len(dropTrailingSpace(nil)) != 0 {
		t.Fatal("dropTrailingSpace wrong")
	}
}
