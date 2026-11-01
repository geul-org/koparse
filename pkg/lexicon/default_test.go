//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what TestDefault — 내장 목록이 읽히고 대표 항목(을, 하는, 자리, Q5, 동음이의어 문서)이 있다
package lexicon

import "testing"

func TestDefault(t *testing.T) {
	lex, err := Default()
	if err != nil {
		t.Fatal(err)
	}
	forms := map[string]bool{}
	for _, s := range lex.Suffixes {
		forms[s.Form+":"+s.Kind] = true
	}
	if !forms["을:particle"] || !forms["하는:verbal"] || !forms["이다:copula"] || !lex.Property["자리"] ||
		!lex.Proper["Q5"] || !lex.Internal["Q4167410"] || len(lex.Idioms) == 0 || !lex.Allow["만\t으로"] {
		t.Fatal("default lists incomplete")
	}
}
