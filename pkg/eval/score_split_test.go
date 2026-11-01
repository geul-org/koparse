//ff:func feature=eval type=judge control=sequence
//ff:what TestScoreSplit — 정답 문장마다 비교하고 틀린 목록을 정렬한다(출력 없는 문장은 정렬 실패)
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
)

func TestScoreSplit(t *testing.T) {
	outs := map[string]analyze.SentenceOut{"1.1": {Eojeol: []analyze.EojeolOut{{Surface: "가", Split: []analyze.Part{{Form: "가"}}}}}}
	gold := map[string][]GoldSplit{"1.1": {{Seg: "가", Cat: "w"}}, "1.2": {{Seg: "나+는", Cat: "n"}, {Seg: "다", Cat: "w"}}}
	s := ScoreSplit(outs, gold)
	if s.Total != 3 || s.Correct != 1 || s.Unaligned != 2 || s.Mismatches[0] > s.Mismatches[1] {
		t.Fatalf("%+v", s)
	}
}
