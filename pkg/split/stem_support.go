//ff:func feature=split type=rule control=selection
//ff:what stemSupport — 일반: 어간이 GEUL 라벨·별칭이면 지지(후보 수가 적을수록 높게). 비한글 덩어리 어간은 foreign_stem, 모르는 어간은 접미 무게×unknown_scale로
package split

import "github.com/park-jun-woo/toulmin/go/pkg/toulmin"

// stemSupport is the warrant of a split candidate. Evidence is the candidate index.
func stemSupport(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f := factsOf(ctx)
	i := candOf(specs)
	c := f.cands[i]
	switch {
	case !hasHangul(c.Stem):
		return f.cfg.ForeignStem, i
	case f.stemCount[i] > 0:
		return knownVerdict(f.stemCount[i], f.cfg.KnownDecay, f.cfg.KnownMin), i
	default:
		return 2*f.cfg.UnknownScale*chainWeight(c.Suffixes) - 1, i
	}
}
