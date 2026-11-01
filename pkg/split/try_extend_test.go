//ff:func feature=split type=builder control=sequence
//ff:what TestTryExtend — 끝이 안 맞음, 빈 어간, 서술격 뒤 사슬, 이미 본 후보는 건너뛴다
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestTryExtend(t *testing.T) {
	lex := &lexicon.Lexicon{}
	eun := lexicon.Suffix{Form: "은", Kind: "particle", Slot: 5}
	ida := lexicon.Suffix{Form: "이다", Kind: "copula", Slot: 9}
	seen := map[string]bool{}
	if out := tryExtend(lex, "책은", ida, nil, 1, nil, seen); out != nil {
		t.Fatal("suffix does not end rest")
	}
	if out := tryExtend(lex, "은", eun, nil, 1, nil, seen); out != nil {
		t.Fatal("empty stem")
	}
	if out := tryExtend(lex, "학생이다", ida, []lexicon.Suffix{eun}, 1, nil, seen); out != nil {
		t.Fatal("copula must be last")
	}
	out := tryExtend(lex, "책은", eun, nil, 1, nil, seen)
	if len(out) != 1 || out[0].Key() != "책|은:particle" {
		t.Fatal(out)
	}
	if again := tryExtend(lex, "책은", eun, nil, 1, nil, seen); again != nil {
		t.Fatal("seen candidate added twice")
	}
}
