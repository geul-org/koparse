//ff:func feature=noun type=rule control=sequence
//ff:what TestClassMismatch — 고유 분류 후보만 +1
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestClassMismatch(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{Proper: true}, {}}})
	a, _ := classMismatch(ctx, toulmin.Specs{nounSpec{Cand: 0}})
	b, _ := classMismatch(ctx, toulmin.Specs{nounSpec{Cand: 1}})
	if a != 1 || b != -1 {
		t.Fatal(a, b)
	}
}
