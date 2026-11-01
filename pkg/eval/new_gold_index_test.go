//ff:func feature=eval type=builder control=sequence
//ff:what TestNewGoldIndex — 문장별·표본별 색인과 손 판정
package eval

import "testing"

func TestNewGoldIndex(t *testing.T) {
	ix := newGoldIndex([]GoldNoun{{Key: "1.2", Noun: "물", Good: []string{"Q283"}}}, map[string]string{"x": Correct})
	if !ix.byKey["1.2\t물"].good["Q283"] || !ix.byItem["1\t물"].good["Q283"] || ix.extra["x"] != Correct {
		t.Fatal(ix)
	}
}
