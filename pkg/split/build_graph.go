//ff:func feature=split type=builder control=iteration dimension=1
//ff:what buildGraph — 어절 하나의 TAN 그래프: 후보마다 warrant, 분리 후보마다 반박 5개, 따옴표 예외 1개
package split

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// buildGraph builds the split graph of n candidates (index 0 = whole).
func buildGraph(n int, cfg config.Split) *toulmin.Graph {
	g := toulmin.NewGraph("split")
	g.Rule(wholeSupport).Qualifier(cfg.QWhole)
	quoted := g.Except(quotedToken)
	for i := 1; i < n; i++ {
		spec := candSpec{Idx: i}
		w := g.Rule(stemSupport).With(spec).Qualifier(cfg.QStem)
		g.Counter(wholeKnown).With(spec).Attacks(w, cfg.ForceWholeKnown)
		g.Counter(batchimMismatch).With(spec).Attacks(w, cfg.ForceBatchim)
		g.Counter(chainInvalid).With(spec).Attacks(w, cfg.ForceChain)
		g.Counter(shortStem).With(spec).Attacks(w, cfg.ForceShortStem)
		g.Counter(shorterKnown).With(spec).Attacks(w, cfg.ForceShorter)
		quoted.Attacks(w, cfg.ForceQuoted)
	}
	return g
}
