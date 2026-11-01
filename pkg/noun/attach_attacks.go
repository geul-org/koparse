//ff:func feature=noun type=builder control=iteration dimension=1
//ff:what attachAttacks — 성질·관용 예외는 모든 warrant를, 분류 불일치 반박은 R1·R4 warrant만 공격
package noun

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

// attachAttacks wires the exceptions to every warrant of a candidate and the
// class-mismatch rebuttal to its weak (R1, R4) warrants.
func attachAttacks(all, weak []*toulmin.Rule, prop, idiom, mismatch *toulmin.Rule, cfg config.Noun) {
	for _, w := range all {
		prop.Attacks(w, cfg.ForceProperty)
		idiom.Attacks(w, cfg.ForceIdiom)
	}
	for _, w := range weak {
		mismatch.Attacks(w, cfg.ForceClass)
	}
}
