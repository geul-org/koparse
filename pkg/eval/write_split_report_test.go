//ff:func feature=eval type=builder control=sequence
//ff:what TestWriteSplitReport — 전체 줄과 범주 줄
package eval

import (
	"strings"
	"testing"
)

func TestWriteSplitReport(t *testing.T) {
	var b strings.Builder
	writeSplitReport(&b, SplitScore{Total: 2, Correct: 1, ByCat: map[string][2]int{"n": {1, 2}}})
	if !strings.Contains(b.String(), "분리 정확도: 1/2 = 50.0%") || !strings.Contains(b.String(), "n") {
		t.Fatal(b.String())
	}
}
