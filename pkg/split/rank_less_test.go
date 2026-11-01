//ff:func feature=split type=judge control=sequence
//ff:what TestRankLess — verdict, 접미 수, 어간 길이 순
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestRankLess(t *testing.T) {
	one := []lexicon.Suffix{{Form: "은"}}
	hi, lo := Scored{Verdict: 0.5}, Scored{Verdict: 0.1}
	whole := Scored{Candidate: Candidate{Stem: "가나"}}
	split := Scored{Candidate: Candidate{Stem: "가", Suffixes: one}}
	long := Scored{Candidate: Candidate{Stem: "가나", Suffixes: one}}
	if !rankLess(hi, lo) || rankLess(lo, hi) || !rankLess(whole, split) || !rankLess(long, split) || rankLess(split, long) {
		t.Fatal("rankLess")
	}
}
