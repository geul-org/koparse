//ff:func feature=eval type=judge control=sequence
//ff:what TestScoreSplitSentence — 맞음, 틀림(미정), 정렬 실패를 센다
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
)

func TestScoreSplitSentence(t *testing.T) {
	out := analyze.SentenceOut{Eojeol: []analyze.EojeolOut{
		{Surface: "고양이는", Split: []analyze.Part{{Form: "고양이"}, {Form: "는"}}, Status: "adopted"},
		{Surface: "작다", Split: []analyze.Part{{Form: "작"}, {Form: "다"}}, Status: "undecided"},
	}}
	rows := []GoldSplit{{Seg: "고양이+는", Cat: "n"}, {Seg: "작다", Cat: "w"}, {Seg: "없음", Cat: "w"}}
	s := SplitScore{ByCat: map[string][2]int{}}
	scoreSplitSentence(&s, "1.1", out, rows)
	if s.Total != 3 || s.Correct != 1 || s.Undecided != 1 || s.Unaligned != 1 || len(s.Mismatches) != 2 || s.ByCat["w"] != [2]int{0, 2} {
		t.Fatalf("%+v", s)
	}
}
