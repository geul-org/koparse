//ff:func feature=stats type=helper control=sequence
//ff:what TestLastSyllable — 마지막 룬, 빈 문자열은 0
package stats

import "testing"

func TestLastSyllable(t *testing.T) {
	if lastSyllable("먹었다") != '다' || lastSyllable("") != 0 {
		t.Fatal("lastSyllable")
	}
}
