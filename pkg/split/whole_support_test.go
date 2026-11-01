//ff:func feature=split type=rule control=sequence
//ff:what TestWholeSupport — 통째가 라벨이면 후보 수 판정, 아니면 사전값
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestWholeSupport(t *testing.T) {
	cfg := config.Split{KnownDecay: 0.1, KnownMin: 0.4, WholePrior: -0.3}
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, &facts{stemCount: []int{1}, cfg: cfg})
	if v, ev := wholeSupport(ctx, nil); v != 1 || ev != 0 {
		t.Fatal(v, ev)
	}
	ctx.Set(factsKey, &facts{stemCount: []int{0}, cfg: cfg})
	if v, _ := wholeSupport(ctx, nil); v != -0.3 {
		t.Fatal(v)
	}
}
