//ff:func feature=noun type=builder control=sequence
//ff:what TestMakeUnit — 끝 단위만 접미와 지배어를 가진다
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestMakeUnit(t *testing.T) {
	w := []Word{{Surface: "자리를", Stem: "자리", Chain: []lexicon.Suffix{{Form: "를", Kind: "particle"}}}, {Surface: "잡고"}}
	fin := makeUnit(w, 0, 0, "자리", "자리", lookupResult{}, false, true)
	mid := makeUnit(w, 0, 0, "자", "자", lookupResult{}, true, false)
	if fin.Chain[0] != "를" || fin.Governor != "잡고" || mid.Chain != nil || mid.Governor != "" || !mid.Piece {
		t.Fatal(fin, mid)
	}
}
