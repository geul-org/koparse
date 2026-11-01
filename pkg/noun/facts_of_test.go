//ff:func feature=noun type=helper control=sequence
//ff:what TestFactsOf — 재료와 후보 번호, Spec이 없으면 −1
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestFactsOf(t *testing.T) {
	f := &nounFacts{property: true}
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, f)
	got, i := factsOf(ctx, toulmin.Specs{nounSpec{Cand: 2}})
	_, j := factsOf(ctx, nil)
	if got != f || i != 2 || j != -1 {
		t.Fatal(got, i, j)
	}
}
