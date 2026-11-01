//ff:func feature=noun type=model control=sequence
//ff:what TestNounSpecSpecName — Spec 이름은 qid
package noun

import "testing"

func TestNounSpecSpecName(t *testing.T) {
	if (nounSpec{}).SpecName() != "qid" {
		t.Fatal("name")
	}
}
