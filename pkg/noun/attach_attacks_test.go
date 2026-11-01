//ff:func feature=noun type=builder control=sequence
//ff:what TestAttachAttacks — 예외는 모든 warrant를 0으로, 반박은 약한 warrant만 깎는다
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestAttachAttacks(t *testing.T) {
	cfg := config.Noun{ForceProperty: 1, ForceIdiom: 1, ForceClass: 1}
	g := toulmin.NewGraph("t")
	prop, idiom := g.Except(propertyFrame), g.Except(idiomFrame)
	spec := nounSpec{Cand: 0}
	strong := g.Rule(r2Topic).With(spec)
	weak := g.Rule(r1Unique).With(spec)
	attachAttacks([]*toulmin.Rule{strong, weak}, []*toulmin.Rule{weak}, prop, idiom, g.Counter(classMismatch).With(spec), cfg)
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{Topic: true, Proper: true}}})
	res, _ := g.Evaluate(ctx)
	if len(res) != 2 || res[0].Verdict != 1 || res[1].Verdict != 0 {
		t.Fatal(res)
	}
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{Topic: true}}, idiom: true})
	res, _ = g.Evaluate(ctx)
	if res[0].Applicability != 0 || res[1].Applicability != 0 {
		t.Fatal(res)
	}
}
