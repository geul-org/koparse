//ff:func feature=noun type=rule control=sequence
//ff:what idiomFrame — 예외: 관용 틀(X를 N으로 삼다, 자리 잡다 …)에 든 명사면 모든 지지를 끊는다
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// idiomFrame gates every warrant of a noun inside an idiom frame.
func idiomFrame(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f, _ := factsOf(ctx, specs)
	return verdictIf(f.idiom), nil
}
