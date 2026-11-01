//ff:func feature=stats type=builder control=sequence
//ff:what TestAppendJ — 빈 조사는 붙이지 않는다
package stats

import "testing"

func TestAppendJ(t *testing.T) {
	toks := appendJ(appendJ(nil, "", 0, 0), "에서는", 2, 1)
	if len(toks) != 1 || toks[0] != (Token{Text: "에서는", Tag: "J", Eojeol: 2, Clause: 1, Unit: -1}) {
		t.Fatal(toks)
	}
}
