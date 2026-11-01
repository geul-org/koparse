//ff:func feature=split type=rule control=sequence
//ff:what wholeSupport — 일반: 통째 후보. 어절 전체가 라벨·별칭이면 후보 수로, 아니면 사전값(whole_prior)
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// wholeSupport is the warrant of the whole-eojeol candidate (index 0).
// Evidence is the candidate index.
func wholeSupport(ctx toulmin.Context, _ toulmin.Specs) (float64, any) {
	f := factsOf(ctx)
	if n := f.stemCount[0]; n > 0 {
		return knownVerdict(n, f.cfg.KnownDecay, f.cfg.KnownMin), 0
	}
	return f.cfg.WholePrior, 0
}
