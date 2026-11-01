//ff:func feature=eval type=builder control=iteration dimension=1
//ff:what TestWriteNounReport — 정밀도·재현율·버림 비율·버림 사유·규칙 줄
package eval

import (
	"strings"
	"testing"
)

func TestWriteNounReport(t *testing.T) {
	var b strings.Builder
	writeNounReport(&b, NounScore{Nouns: 4, Adopted: 2, Dropped: 2, Judged: map[string]int{Correct: 1, Wrong: 1},
		DroppedReason: map[string]int{"성질": 1, "문턱 미달": 1}, ByRule: map[string][3]int{"R2 주제": {1, 0, 0}}, GoldGood: 2, Recalled: 1})
	out := b.String()
	for _, want := range []string{"정밀도: 1/2 = 50.0%", "재현율: 1/2", "성질·관용 버림 비율: 1/4", "버림 사유 문턱 미달", "규칙 R2 주제"} {
		if !strings.Contains(out, want) {
			t.Errorf("missing %q in %s", want, out)
		}
	}
}
