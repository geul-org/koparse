//ff:func feature=eval type=builder control=sequence
//ff:what WriteReport — 분리·명사 요약과 fixture 놓침 수를 쓴다
package eval

import (
	"fmt"
	"io"
)

// WriteReport prints the whole evaluation summary.
func WriteReport(w io.Writer, sp SplitScore, ns NounScore, misses int) {
	writeSplitReport(w, sp)
	writeNounReport(w, ns)
	fmt.Fprintf(w, "fixture에 없어 \"후보 없음\"으로 답한 표면: %d\n", misses)
}
