//ff:func feature=split type=builder control=iteration dimension=1
//ff:what TestExtend — 깊이 0이면 그대로, 깊이만큼만 사슬을 늘린다
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestExtend(t *testing.T) {
	lex, _ := lexicon.Default()
	if got := extend(lex, "선수들은", nil, 0, nil, map[string]bool{}); got != nil {
		t.Fatal(got)
	}
	one := extend(lex, "선수들은", nil, 1, nil, map[string]bool{})
	for _, c := range one {
		if len(c.Suffixes) != 1 {
			t.Errorf("depth 1 gave %s", c.Key())
		}
	}
	if len(one) == 0 {
		t.Fatal("no candidates")
	}
}
