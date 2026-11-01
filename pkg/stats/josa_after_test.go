//ff:func feature=stats type=helper control=sequence
//ff:what TestJosaAfter — 같은 어절의 바로 뒤 J만, 아니면 ∅
package stats

import "testing"

func TestJosaAfter(t *testing.T) {
	toks := []Token{{Tag: "N", Eojeol: 0}, {Tag: "J", Text: "는", Eojeol: 0}, {Tag: "N", Eojeol: 1}, {Tag: "J", Text: "의", Eojeol: 2}}
	if josaAfter(toks, 0) != "는" || josaAfter(toks, 2) != "∅" || josaAfter(toks, 3) != "∅" {
		t.Fatal("josaAfter")
	}
}
