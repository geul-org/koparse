//ff:func feature=eval type=helper control=sequence
//ff:what TestSplitList — 빈 문자열은 nil, 쉼표로 나눈다
package eval

import "testing"

func TestSplitList(t *testing.T) {
	if splitList(" ") != nil || len(splitList("Q1,Q2")) != 2 {
		t.Fatal("splitList")
	}
}
