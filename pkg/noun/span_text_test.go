//ff:func feature=noun type=helper control=sequence
//ff:what TestSpanText — 어간을 구분자로 잇는다
package noun

import "testing"

func TestSpanText(t *testing.T) {
	w := []Word{{Stem: "한강"}, {Stem: "연안"}, {Stem: "지역"}}
	if spanText(w, 0, 1, " ") != "한강 연안" || spanText(w, 1, 2, "") != "연안지역" {
		t.Fatal("spanText")
	}
}
