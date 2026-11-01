//ff:func feature=eval type=builder control=iteration dimension=1
//ff:what writeNounReport — 명사 식별 요약(상태별 수, 정밀도, 재현율, 버림 비율, 규칙별 판정)을 쓴다
package eval

import (
	"fmt"
	"io"
)

// writeNounReport prints the noun identification summary. Precision counts
// only judged adoptions; unjudged ones are reported separately.
func writeNounReport(w io.Writer, s NounScore) {
	c, x, u := s.Judged[Correct], s.Judged[Wrong], s.Judged[Unjudged]
	fmt.Fprintf(w, "명사: 전체 %d, 채택 %d, 버림 %d, 미정 %d\n", s.Nouns, s.Adopted, s.Dropped, s.Undecided)
	fmt.Fprintf(w, "  정밀도: %s (미판정 %d)\n", pct(c, c+x), u)
	fmt.Fprintf(w, "  재현율: %s (맞는 QID가 있는 정답 명사 기준)\n", pct(s.Recalled, s.GoldGood))
	fmt.Fprintf(w, "  버림·미정 비율: %s\n", pct(s.Dropped+s.Undecided, s.Nouns))
	fmt.Fprintf(w, "  성질·관용 버림 비율: %s\n", pct(s.DroppedReason["성질"]+s.DroppedReason["관용"], s.Nouns))
	for _, r := range sortedMapKeys(s.DroppedReason) {
		fmt.Fprintf(w, "  버림 사유 %-8s %d\n", r, s.DroppedReason[r])
	}
	for _, r := range sortedMapKeys(s.ByRule) {
		b := s.ByRule[r]
		fmt.Fprintf(w, "  규칙 %-14s 정밀도 %s (미판정 %d)\n", r, pct(b[0], b[0]+b[1]), b[2])
	}
}
