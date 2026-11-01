//ff:func feature=lexicon type=parser control=sequence
//ff:what TestParseWordList — 첫 필드만 집합으로
package lexicon

import "testing"

func TestParseWordList(t *testing.T) {
	set := parseWordList("# c\n자리\n기준\textra\n")
	if len(set) != 2 || !set["자리"] || !set["기준"] {
		t.Fatalf("got %v", set)
	}
}
