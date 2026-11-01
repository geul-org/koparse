//ff:func feature=split type=rule control=sequence
//ff:what TestShortStem — 한 음절 한글 어간만 +1(두 음절, 한 글자 영문은 −1)
package split

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestShortStem(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, &facts{cands: []Candidate{{Stem: "물"}, {Stem: "물질"}, {Stem: "X"}}})
	a, _ := shortStem(ctx, toulmin.Specs{candSpec{Idx: 0}})
	b, _ := shortStem(ctx, toulmin.Specs{candSpec{Idx: 1}})
	c, _ := shortStem(ctx, toulmin.Specs{candSpec{Idx: 2}})
	if a != 1 || b != -1 || c != -1 {
		t.Fatal(a, b, c)
	}
}
