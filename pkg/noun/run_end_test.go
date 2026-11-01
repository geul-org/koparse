//ff:func feature=noun type=helper control=sequence
//ff:what TestRunEnd — 맨 어절은 이어 붙이고, 접미·쉼표·끝에서 멈춘다
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestRunEnd(t *testing.T) {
	p := []lexicon.Suffix{{Form: "에"}}
	w := []Word{{Stem: "한강"}, {Stem: "연안", Chain: p}, {Stem: "a", Break: true}, {Stem: "b"}, {Stem: "c"}}
	if runEnd(w, 0) != 1 || runEnd(w, 1) != 1 || runEnd(w, 2) != 2 || runEnd(w, 3) != 4 {
		t.Fatal(runEnd(w, 0), runEnd(w, 2), runEnd(w, 3))
	}
}
