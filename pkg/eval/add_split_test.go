//ff:func feature=eval type=judge control=sequence
//ff:what TestAddSplit — 맞은 미정은 따로 세고, 틀리면 목록에 남긴다
package eval

import "testing"

func TestAddSplit(t *testing.T) {
	s := SplitScore{ByCat: map[string][2]int{}}
	addSplit(&s, "1.1", GoldSplit{Seg: "가+는", Cat: "c"}, "가+는", true)
	addSplit(&s, "1.1", GoldSplit{Seg: "나", Cat: "c"}, "나+다", false)
	if s.Total != 2 || s.Correct != 1 || s.Undecided != 1 || s.UndecidedCorrect != 1 || s.ByCat["c"] != [2]int{1, 2} || s.Mismatches[0] != "1.1\t나\t나+다\tc" {
		t.Fatalf("%+v", s)
	}
}
