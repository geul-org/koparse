//ff:func feature=stats type=builder control=sequence
//ff:what TestStemToken — 접미 없는 서술어는 V, 그 밖은 M
package stats

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestStemToken(t *testing.T) {
	v := stemToken(noun.Word{Surface: "달린다", Stem: "달린다"}, 1, 0)
	m := stemToken(noun.Word{Surface: "km2로", Stem: "km2", Chain: []lexicon.Suffix{{Form: "로"}}}, 2, 0)
	if v.Tag != "V" || v.Eojeol != 1 || m.Tag != "M" || m.Text != "km2" {
		t.Fatal(v, m)
	}
}
