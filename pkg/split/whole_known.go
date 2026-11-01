//ff:func feature=split type=rule control=sequence
//ff:what wholeKnown — 반박: 어절 전체가 라벨·별칭이면 분리 후보를 공격(가수 → 가+수)
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// wholeKnown rebuts a split candidate when the whole eojeol is a known label.
func wholeKnown(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	if factsOf(ctx).stemCount[0] > 0 {
		return 1, nil
	}
	return -1, nil
}
