//ff:func feature=split type=rule control=sequence
//ff:what quotedToken — 예외: 따옴표 안 어절이면 분리 규칙을 끊는다(적용도 게이트)
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// quotedToken gates every split candidate of a quoted eojeol.
func quotedToken(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	if factsOf(ctx).quoted {
		return 1, nil
	}
	return -1, nil
}
