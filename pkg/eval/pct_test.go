//ff:func feature=eval type=helper control=sequence
//ff:what TestPct — 분모 0은 "-", 아니면 소수 한 자리 백분율
package eval

import "testing"

func TestPct(t *testing.T) {
	if pct(1, 0) != "1/0 = -" || pct(1, 3) != "1/3 = 33.3%" {
		t.Fatal(pct(1, 3))
	}
}
