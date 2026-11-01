//ff:func feature=stats type=helper control=sequence
//ff:what TestFindTag — 앞으로·뒤로 찾고, 없으면 −1
package stats

import "testing"

func TestFindTag(t *testing.T) {
	toks := []Token{{Tag: "N"}, {Tag: "V"}, {Tag: "N"}, {Tag: "V"}}
	if findTag(toks, 0, 1, "V") != 1 || findTag(toks, 3, -1, "N") != 2 || findTag(toks, 0, 1, "J") != -1 || findTag(toks, 5, 1, "N") != -1 {
		t.Fatal("findTag")
	}
}
