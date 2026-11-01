//ff:func feature=noun type=rule control=sequence
//ff:what classMismatch — 반박: 후보가 고유 분류(인명·이름·글자·작품 등)면 R1·R4 지지를 공격
package noun

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// classMismatch rebuts R1 and R4 support for a candidate of a proper class.
func classMismatch(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f, i := factsOf(ctx, specs)
	return verdictIf(f.cands[i].Proper), nil
}
