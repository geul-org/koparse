//ff:func feature=split type=helper control=iteration dimension=1
//ff:what TestBatchimMismatchAt — 사과→사+과 어긋남, 명으+로는 으로로 봄, 맞는 사슬은 빈 문자열
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestBatchimMismatchAt(t *testing.T) {
	gwa := lexicon.Suffix{Form: "과", Cond: "C"}
	ro := lexicon.Suffix{Form: "로", Cond: "VL"}
	eul := lexicon.Suffix{Form: "을", Cond: "C"}
	cases := []struct {
		c    Candidate
		want string
	}{
		{Candidate{Stem: "사", Suffixes: []lexicon.Suffix{gwa}}, "과"},
		{Candidate{Stem: "명으", Suffixes: []lexicon.Suffix{ro}}, "로"},
		{Candidate{Stem: "물", Suffixes: []lexicon.Suffix{ro}}, ""},
		{Candidate{Stem: "책", Suffixes: []lexicon.Suffix{gwa, eul}}, "을"},
	}
	for _, c := range cases {
		if got := batchimMismatchAt(c.c); got != c.want {
			t.Errorf("%s = %q, want %q", c.c.Key(), got, c.want)
		}
	}
}
