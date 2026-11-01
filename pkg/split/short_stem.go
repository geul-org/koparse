//ff:func feature=split type=rule control=sequence
//ff:what shortStem — 반박(약한 세기): 한 음절 한글 어간이면 공격
package split

import (
	"github.com/geul-org/koparse/pkg/hangul"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// shortStem weakly rebuts a split whose stem is a single Hangul syllable.
func shortStem(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	stem := factsOf(ctx).cands[candOf(specs)].Stem
	if hangul.Count(stem) == 1 && hangul.IsAllHangul(stem) {
		return 1, nil
	}
	return -1, nil
}
