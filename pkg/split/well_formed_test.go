//ff:func feature=split type=helper control=sequence
//ff:what TestWellFormed — 사슬 순서와 받침이 다 맞아야 참
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestWellFormed(t *testing.T) {
	f := &facts{lex: &lexicon.Lexicon{}}
	ro := lexicon.Suffix{Form: "로", Cond: "VL", Slot: 2}
	eul := lexicon.Suffix{Form: "을", Cond: "C", Slot: 4}
	if !wellFormed(f, Candidate{Stem: "물", Suffixes: []lexicon.Suffix{ro}}) ||
		wellFormed(f, Candidate{Stem: "물", Suffixes: []lexicon.Suffix{eul, ro}}) ||
		wellFormed(f, Candidate{Stem: "사", Suffixes: []lexicon.Suffix{eul}}) {
		t.Fatal("wellFormed")
	}
}
