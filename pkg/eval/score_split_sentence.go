//ff:func feature=eval type=judge control=iteration dimension=1
//ff:what scoreSplitSentence — 문장 하나: 정답 어절마다 같은 표면의 출력 어절을 찾아 경계를 비교한다
package eval

import (
	"strings"

	"github.com/geul-org/koparse/pkg/analyze"
)

// scoreSplitSentence adds one sentence to s. A gold eojeol without an output
// eojeol of the same surface counts as wrong (unaligned).
func scoreSplitSentence(s *SplitScore, key string, out analyze.SentenceOut, rows []GoldSplit) {
	used := make([]bool, len(out.Eojeol))
	for _, g := range rows {
		got, undecided := "(정렬 실패)", false
		if i := findEojeol(out.Eojeol, used, strings.ReplaceAll(g.Seg, "+", "")); i >= 0 {
			used[i] = true
			got, undecided = segOf(out.Eojeol[i]), out.Eojeol[i].Status != "adopted"
		} else {
			s.Unaligned++
		}
		addSplit(s, key, g, got, undecided)
	}
}
