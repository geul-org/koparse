//ff:func feature=noun type=builder control=sequence
//ff:what TestBuildNounGraph — 후보 둘: 살아 있는 warrant만 결과로, 고유 분류 반박은 R1을 깎고 R2는 그대로
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestBuildNounGraph(t *testing.T) {
	cfg := config.Default().Noun
	ctx := toulmin.NewContext()
	ctx.Set(nounFactsKey, &nounFacts{cands: []candInfo{{Topic: true, Proper: true}, {}}})
	res, err := buildNounGraph(2, cfg).Evaluate(ctx)
	if err != nil || len(res) != 1 || res[0].Evidence.(hit).Rule != RuleR2 || res[0].Verdict != 2*cfg.QR2-1 {
		t.Fatal(res, err)
	}
}
