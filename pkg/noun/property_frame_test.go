//ff:func feature=noun type=rule control=sequence
//ff:what TestPropertyFrame — 성질·틀 명사면 +1
package noun

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestPropertyFrame(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{property: true})
	a, _ := propertyFrame(ctx, nil)
	ctx.Set(nounFactsKey, &nounFacts{})
	b, _ := propertyFrame(ctx, nil)
	if a != 1 || b != -1 {
		t.Fatal(a, b)
	}
}
