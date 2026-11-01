//ff:func feature=eval type=builder control=iteration dimension=1
//ff:what writeSplitReport — 분리 정확도 요약(전체, 미정, 정렬 실패, 범주별)을 쓴다
package eval

import (
	"fmt"
	"io"
)

// writeSplitReport prints the split accuracy summary.
func writeSplitReport(w io.Writer, s SplitScore) {
	fmt.Fprintf(w, "분리 정확도: %s (미정 %d, 그중 맞음 %d, 정렬 실패 %d)\n", pct(s.Correct, s.Total), s.Undecided, s.UndecidedCorrect, s.Unaligned)
	for _, c := range sortedMapKeys(s.ByCat) {
		fmt.Fprintf(w, "  %-14s %s\n", c, pct(s.ByCat[c][0], s.ByCat[c][1]))
	}
}
