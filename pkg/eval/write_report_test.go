//ff:func feature=eval type=builder control=sequence
//ff:what TestWriteReport — 분리·명사 요약과 놓침 수
package eval

import (
	"strings"
	"testing"
)

func TestWriteReport(t *testing.T) {
	var b strings.Builder
	WriteReport(&b, SplitScore{ByCat: map[string][2]int{}}, NounScore{Judged: map[string]int{}}, 7)
	if !strings.Contains(b.String(), "분리 정확도") || !strings.Contains(b.String(), "명사:") || !strings.Contains(b.String(), ": 7") {
		t.Fatal(b.String())
	}
}
