//ff:func feature=noun type=helper control=sequence
//ff:what TestVerdictIf — 참 +1, 거짓 −1
package noun

import "testing"

func TestVerdictIf(t *testing.T) {
	if verdictIf(true) != 1 || verdictIf(false) != -1 {
		t.Fatal("verdictIf")
	}
}
