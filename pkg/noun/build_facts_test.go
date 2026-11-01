//ff:func feature=noun type=builder control=sequence
//ff:what TestBuildFacts — 주제·진술·고유 분류·R4a·성질·관용 표시를 채운다(빈 주제는 일치하지 않음)
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestBuildFacts(t *testing.T) {
	lex, _ := lexicon.Default()
	u := Unit{Form: "자리", Chain: []string{}, Governor: "잡고",
		Cands: []knowledge.Candidate{{QID: "Q1"}, {QID: "Q2"}, {QID: "Q3"}},
		Infos: []knowledge.EntityInfo{{}, {P31: []string{"Q5"}}, {P31: []string{"Q5"}}}}
	f := buildFacts(u, DocContext{Topic: "Q1", Statements: map[string]bool{"Q2": true}}, lex)
	c := f.cands
	if !c[0].Topic || !c[1].Statement || !c[1].Proper || c[0].Proper || !c[0].General || !f.property || !f.idiom {
		t.Fatalf("%+v", f)
	}
	if g := buildFacts(Unit{Cands: []knowledge.Candidate{{QID: ""}}, Infos: []knowledge.EntityInfo{{}}}, DocContext{}, lex); g.cands[0].Topic {
		t.Fatal("empty topic matched")
	}
}
