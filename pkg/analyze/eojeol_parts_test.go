//ff:func feature=analyze type=builder control=sequence
//ff:what TestEojeolParts — 접미가 붙거나 덮인 어간은 noun, 아니면 word, 접미는 종류대로
package analyze

import (
	"reflect"
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
)

func TestEojeolParts(t *testing.T) {
	got := eojeolParts(noun.Word{Stem: "고기", Chain: []lexicon.Suffix{{Form: "를", Kind: "particle"}}}, false)
	if !reflect.DeepEqual(got, []Part{{"고기", "noun"}, {"를", "particle"}}) {
		t.Fatal(got)
	}
	if eojeolParts(noun.Word{Stem: "먹는다"}, false)[0].Type != "word" || eojeolParts(noun.Word{Stem: "한강"}, true)[0].Type != "noun" {
		t.Fatal("types")
	}
}
