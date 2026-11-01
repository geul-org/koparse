//ff:func feature=noun type=rule control=sequence
//ff:what propertyFrame — 예외: 성질·틀 명사(자리, 기준, 사이 …)면 모든 지지를 끊는다(개체로 쓰이지 않음)
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// propertyFrame gates every warrant of a property/frame noun.
func propertyFrame(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f, _ := factsOf(ctx, specs)
	return verdictIf(f.property), nil
}
