//ff:func feature=text type=helper control=sequence
//ff:what TestHeadBefore — 마지막 낱말에서 문장부호를 떼고, 빈 입력은 빈 문자열
package text

import "testing"

func TestHeadBefore(t *testing.T) {
	if headBefore([]rune("커피, 서울특별시 ")) != "서울특별시" || headBefore([]rune("“광합성")) != "광합성" || headBefore(nil) != "" {
		t.Fatal("headBefore wrong")
	}
}
