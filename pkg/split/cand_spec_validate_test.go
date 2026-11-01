//ff:func feature=split type=model control=sequence
//ff:what TestCandSpecValidate — 음수 번호는 오류
package split

import "testing"

func TestCandSpecValidate(t *testing.T) {
	if (candSpec{Idx: 0}).Validate() != nil || (candSpec{Idx: -1}).Validate() == nil {
		t.Fatal("validate")
	}
}
