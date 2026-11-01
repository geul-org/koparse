//ff:func feature=text type=helper control=sequence
//ff:what TestIsSentenceEnd — 줄바꿈, 공백·끝 앞의 . ? !는 끝, 605.25의 점은 아님
package text

import "testing"

func TestIsSentenceEnd(t *testing.T) {
	rs := []rune("가.나 605.25? 다\n라!")
	if isSentenceEnd(rs, 1) || isSentenceEnd(rs, 7) || !isSentenceEnd(rs, 10) || !isSentenceEnd(rs, 13) || !isSentenceEnd(rs, 15) || isSentenceEnd(rs, 0) {
		t.Fatal("isSentenceEnd wrong")
	}
}
