//ff:func feature=split type=helper control=sequence
//ff:what factsOf — toulmin 문맥에서 판정 재료를 꺼낸다
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// factsOf returns the *facts stored in ctx.
func factsOf(ctx toulmin.Context) *facts {
	v, _ := ctx.Get(factsKey)
	return v.(*facts)
}
