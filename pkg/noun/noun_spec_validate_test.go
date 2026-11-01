//ff:func feature=noun type=model control=sequence
//ff:what TestNounSpecValidate — 음수 번호는 오류
package noun

import "testing"

func TestNounSpecValidate(t *testing.T) {
	if (nounSpec{Cand: 1}).Validate() != nil || (nounSpec{Cand: -1}).Validate() == nil {
		t.Fatal("validate")
	}
}
