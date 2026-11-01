//ff:func feature=split type=rule control=sequence
//ff:what batchimMismatch — 반박: 조사가 앞 음절 받침과 어긋나면 공격(사과 → 사+과)
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// batchimMismatch rebuts a split whose suffix allomorph does not fit the
// preceding sound. Evidence is the offending suffix.
func batchimMismatch(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	bad := batchimMismatchAt(factsOf(ctx).cands[candOf(specs)])
	if bad != "" {
		return 1, bad
	}
	return -1, nil
}
