//ff:func feature=eval type=judge control=iteration dimension=1
//ff:what ScoreSplit — 정답이 있는 문장마다 분리 경계를 비교해 정확도를 모은다(출력 없는 정답 문장은 모두 정렬 실패)
package eval

import (
	"sort"

	"github.com/geul-org/koparse/pkg/analyze"
)

// ScoreSplit compares the output splits with the gold splits.
func ScoreSplit(outs map[string]analyze.SentenceOut, gold map[string][]GoldSplit) SplitScore {
	s := SplitScore{ByCat: map[string][2]int{}}
	for key, rows := range gold {
		scoreSplitSentence(&s, key, outs[key], rows)
	}
	sort.Strings(s.Mismatches)
	return s
}
