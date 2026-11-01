//ff:func feature=split type=rule control=iteration dimension=1
//ff:what shorterKnown — 반박: 모르는 어간 안에 아는 더 짧은 어간이 있고 그 분리가 사슬·받침에 맞으면 공격(물로+부터 ← 물+로부터)
package split

import (
	"strings"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// shorterKnown rebuts a split whose stem is unknown while another, well-formed
// candidate (valid chain, matching allomorphs) has a known stem that is a
// proper prefix of it: the longer stem still holds a suffix.
func shorterKnown(ctx toulmin.Context, specs toulmin.Specs) (float64, any) {
	f := factsOf(ctx)
	i := candOf(specs)
	if f.stemCount[i] > 0 {
		return -1, nil
	}
	for j, c := range f.cands {
		if f.stemCount[j] > 0 && len(c.Stem) < len(f.cands[i].Stem) && strings.HasPrefix(f.cands[i].Stem, c.Stem) && wellFormed(f, c) {
			return 1, c.Stem
		}
	}
	return -1, nil
}
