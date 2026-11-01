//ff:func feature=split type=judge control=sequence
//ff:what TestJudge — 아는 어간 분리는 채택, 비등하면 미정, 문턱 미달도 미정, 규칙 오류는 오류
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestJudge(t *testing.T) {
	lex, _ := lexicon.Default()
	cfg := config.Default().Split
	cands := []Candidate{{Stem: "책을"}, {Stem: "책", Suffixes: []lexicon.Suffix{{Form: "을", Kind: "particle", Slot: 4, Cond: "C", Weight: 0.9}}}}
	d, err := judge(&facts{cands: cands, stemCount: []int{0, 1}, lex: lex, cfg: cfg})
	if err != nil || d.Status != StatusAdopted || d.Ranked[0].Candidate.Stem != "책" {
		t.Fatal(d, err)
	}
	tie := cfg
	tie.Margin = 10
	if d, _ := judge(&facts{cands: cands, stemCount: []int{0, 1}, lex: lex, cfg: tie}); d.Status != StatusUndecided {
		t.Fatal(d)
	}
	high := cfg
	high.Adopt = 2
	if d, _ := judge(&facts{cands: cands[:1], stemCount: []int{0}, lex: lex, cfg: high}); d.Status != StatusUndecided {
		t.Fatal(d)
	}
	if _, err := judge(&facts{cands: cands, stemCount: []int{0}, lex: lex, cfg: cfg}); err == nil {
		t.Fatal("want error from a panicking rule")
	}
}
