//ff:func feature=stats type=builder control=sequence
//ff:what TestPairsFor — 자기 자신과 M은 빼고 구간·같은 어절·절 경계를 적는다
package stats

import "testing"

func TestPairsFor(t *testing.T) {
	toks := []Token{{Tag: "N", Text: "물", Eojeol: 0}, {Tag: "J", Text: "이", Eojeol: 0}, {Tag: "M", Text: "아주", Eojeol: 1}, {Tag: "V", Text: "맑다", Eojeol: 2, Clause: 1}}
	got := pairsFor(toks, 0)
	if len(got) != 2 || got[0] != (Pair{Other: "이", Tag: "J", Bucket: "+1", SameEojeol: true}) || got[1] != (Pair{Other: "맑다", Tag: "V", Bucket: "+3~4", CrossClause: true}) {
		t.Fatal(got)
	}
}
