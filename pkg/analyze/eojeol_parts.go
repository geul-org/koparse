//ff:func feature=analyze type=builder control=iteration dimension=1
//ff:what eojeolParts — 분리된 어절을 출력 조각으로(어간은 명사 단위에 덮였거나 접미가 붙으면 noun, 아니면 word)
package analyze

import "github.com/geul-org/koparse/pkg/noun"

// eojeolParts returns the output parts of w; covered tells whether a noun
// unit covers the eojeol.
func eojeolParts(w noun.Word, covered bool) []Part {
	t := "word"
	if covered || len(w.Chain) > 0 {
		t = "noun"
	}
	parts := []Part{{Form: w.Stem, Type: t}}
	for _, s := range w.Chain {
		parts = append(parts, Part{Form: s.Form, Type: s.Kind})
	}
	return parts
}
