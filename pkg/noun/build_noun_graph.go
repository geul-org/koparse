//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what buildNounGraph — 명사 하나의 TAN 그래프: (후보, R1~R4b)마다 warrant, 후보마다 분류 불일치 반박, 성질·관용 예외
package noun

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// buildNounGraph builds the noun graph of n QID candidates. Each (candidate,
// general rule) pair is a warrant; the candidate's support is their noisy-OR.
func buildNounGraph(n int, cfg config.Noun) *toulmin.Graph {
	g := toulmin.NewGraph("noun")
	prop := g.Except(propertyFrame)
	idiom := g.Except(idiomFrame)
	for i := 0; i < n; i++ {
		spec := nounSpec{Cand: i}
		ws := []*toulmin.Rule{
			g.Rule(r2Topic).With(spec).Qualifier(cfg.QR2),
			g.Rule(r3Statement).With(spec).Qualifier(cfg.QR3),
			g.Rule(r1Unique).With(spec).Qualifier(cfg.QR1),
			g.Rule(r4aGeneral).With(spec).Qualifier(cfg.QR4A),
			g.Rule(r4bNeighbor).With(spec).Qualifier(cfg.QR4B),
		}
		mismatch := g.Counter(classMismatch).With(spec)
		attachAttacks(ws, ws[2:], prop, idiom, mismatch, cfg)
	}
	return g
}
