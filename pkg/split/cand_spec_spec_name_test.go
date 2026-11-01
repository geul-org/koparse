//ff:func feature=split type=model control=sequence
//ff:what TestCandSpecSpecName — Spec 이름은 cand
package split

import "testing"

func TestCandSpecSpecName(t *testing.T) {
	if (candSpec{}).SpecName() != "cand" {
		t.Fatal("name")
	}
}
