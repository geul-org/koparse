//ff:func feature=split type=helper control=sequence
//ff:what TestKnownVerdict — 후보 하나는 1, 많아지면 내려가되 하한에서 멈춘다
package split

import (
	"math"
	"testing"
)

func TestKnownVerdict(t *testing.T) {
	if knownVerdict(1, 0.1, 0.4) != 1 || math.Abs(knownVerdict(10, 0.1, 0.4)-(1-0.1*math.Log(10))) > 1e-9 || knownVerdict(1e9, 0.1, 0.4) != 0.4 {
		t.Fatal("knownVerdict")
	}
}
