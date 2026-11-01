//ff:func feature=noun type=rule control=sequence
//ff:what TestIdiomFrame — 관용 틀이면 +1
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestIdiomFrame(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{idiom: true})
	a, _ := idiomFrame(ctx, nil)
	ctx.Set(nounFactsKey, &nounFacts{})
	b, _ := idiomFrame(ctx, nil)
	if a != 1 || b != -1 {
		t.Fatal(a, b)
	}
}
