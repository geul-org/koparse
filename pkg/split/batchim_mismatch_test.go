//ff:func feature=split type=rule control=sequence
//ff:what TestBatchimMismatch — 사+과는 +1(근거 과), 책+과는 −1
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestBatchimMismatch(t *testing.T) {
	gwa := []lexicon.Suffix{{Form: "과", Cond: "C"}}
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, &facts{cands: []Candidate{{Stem: "사과"}, {Stem: "사", Suffixes: gwa}, {Stem: "책", Suffixes: gwa}}})
	v1, ev := batchimMismatch(ctx, toulmin.Specs{candSpec{Idx: 1}})
	v2, _ := batchimMismatch(ctx, toulmin.Specs{candSpec{Idx: 2}})
	if v1 != 1 || ev != "과" || v2 != -1 {
		t.Fatal(v1, ev, v2)
	}
}
