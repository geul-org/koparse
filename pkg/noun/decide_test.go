//ff:func feature=noun type=judge control=iteration dimension=1
//ff:what TestDecide — 성질·관용 버림, 문턱 미달, 애매, 채택(+R5 표시)
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
)

func TestDecide(t *testing.T) {
	cfg := config.Noun{Adopt: 0.5, Margin: 0.2}
	two := []candInfo{{QID: "Q1"}, {QID: "Q2"}}
	gated := []candScore{{Miss: 1}, {Miss: 1}}
	cases := []struct {
		u      Unit
		f      nounFacts
		s      []candScore
		status string
		reason string
	}{
		{Unit{Form: "자리", Matched: "자리"}, nounFacts{cands: two, property: true}, gated, StatusDropped, ReasonProperty},
		{Unit{Form: "흉내", Matched: "흉내"}, nounFacts{cands: two, idiom: true}, gated, StatusDropped, ReasonIdiom},
		{Unit{Form: "전통", Matched: "전통"}, nounFacts{cands: two}, []candScore{{Miss: 0.4, Applicability: 1}, {Miss: 1}}, StatusDropped, ReasonBelow},
		{Unit{Form: "수도", Matched: "수도"}, nounFacts{cands: two}, []candScore{{Miss: 0.05, Applicability: 1}, {Miss: 0.1, Applicability: 1}}, StatusUndecided, ReasonAmbiguous},
		{Unit{Form: "발효식품", Matched: "발효 식품"}, nounFacts{cands: two[:1]}, []candScore{{Miss: 0.05, Applicability: 1, Rules: []string{RuleR3}}}, StatusAdopted, ""},
	}
	for _, c := range cases {
		r := decide(c.u, c.f, c.s, cfg)
		if r.Status != c.status || r.Reason != c.reason {
			t.Errorf("%s: %+v", c.u.Form, r)
		}
	}
	r := decide(cases[4].u, cases[4].f, cases[4].s, cfg)
	if r.QID != "Q1" || r.Rules[len(r.Rules)-1] != RuleR5 {
		t.Fatal(r)
	}
}
