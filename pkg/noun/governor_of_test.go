//ff:func feature=noun type=helper control=sequence
//ff:what TestGovernorOf — 자기 서술격·용언화 접미, 아니면 다음 어절, 끝이면 빈 문자열
package noun

import (
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestGovernorOf(t *testing.T) {
	w := []Word{
		{Surface: "근간으로한", Chain: []lexicon.Suffix{{Form: "으로", Kind: "particle"}, {Form: "한", Kind: "verbal"}}},
		{Surface: "자리", Stem: "자리"},
		{Surface: "잡고"},
	}
	if governorOf(w, 0) != "한" || governorOf(w, 1) != "잡고" || governorOf(w, 2) != "" {
		t.Fatal(governorOf(w, 0), governorOf(w, 1))
	}
}
