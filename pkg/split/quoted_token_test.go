//ff:func feature=split type=rule control=sequence
//ff:what TestQuotedToken — 인용 어절만 +1
package split

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestQuotedToken(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, &facts{quoted: true})
	a, _ := quotedToken(ctx, nil)
	ctx.Set(factsKey, &facts{})
	b, _ := quotedToken(ctx, nil)
	if a != 1 || b != -1 {
		t.Fatal(a, b)
	}
}
