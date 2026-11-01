//ff:func feature=split type=helper control=sequence
//ff:what TestCandidateKey — 어간과 접미 형태·종류로 열쇠를 만든다
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestCandidateKey(t *testing.T) {
	c := Candidate{Stem: "선수", Suffixes: []lexicon.Suffix{{Form: "들", Kind: "particle"}, {Form: "은", Kind: "particle"}}}
	if c.Key() != "선수|들:particle|은:particle" || (Candidate{Stem: "가수"}).Key() != "가수" {
		t.Fatal(c.Key())
	}
}
