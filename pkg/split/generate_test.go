//ff:func feature=split type=builder control=iteration dimension=1
//ff:what TestGenerate — 통째 후보가 먼저, 선수들은의 세 분리, 서술격은 맨 끝에만, 빈 어간 없음
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestGenerate(t *testing.T) {
	lex, _ := lexicon.Default()
	keys := map[string]bool{}
	cands := Generate(lex, "선수들은")
	for _, c := range cands {
		keys[c.Key()] = true
	}
	if cands[0].Key() != "선수들은" || !keys["선수들|은:particle"] || !keys["선수|들:particle|은:particle"] {
		t.Fatal(keys)
	}
	for _, c := range Generate(lex, "학생들이다") {
		if len(c.Suffixes) > 1 && c.Suffixes[0].Slot == 9 {
			t.Errorf("copula not last: %s", c.Key())
		}
	}
	if got := Generate(lex, "을"); len(got) != 1 {
		t.Fatalf("a bare suffix has no stem: %v", got)
	}
}
