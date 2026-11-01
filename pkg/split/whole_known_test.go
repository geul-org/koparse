//ff:func feature=split type=rule control=sequence
//ff:what TestWholeKnown — 어절 전체가 라벨일 때만 +1
package split

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWholeKnown(t *testing.T) {
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, &facts{stemCount: []int{2}})
	a, _ := wholeKnown(ctx, nil)
	ctx.Set(factsKey, &facts{stemCount: []int{0}})
	b, _ := wholeKnown(ctx, nil)
	if a != 1 || b != -1 {
		t.Fatal(a, b)
	}
}
