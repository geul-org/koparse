//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what TestFromFS — 네 파일을 읽고, 파일이 없거나 형식이 틀리면 오류
package lexicon

import (
	"maps"
	"testing"
	"testing/fstest"
)

func TestFromFS(t *testing.T) {
	good := fstest.MapFS{
		"d/suffixes.tsv":       {Data: []byte("을\tparticle\t4\tC\t0.9\n")},
		"d/property_nouns.txt": {Data: []byte("자리\n")},
		"d/idioms.tsv":         {Data: []byte("자리\t-\t잡\n")},
		"d/classes.tsv":        {Data: []byte("proper\tQ5\n")},
	}
	lex, err := FromFS(good, "d")
	if err != nil || len(lex.Suffixes) != 1 || !lex.Property["자리"] || len(lex.Idioms) != 1 || !lex.Proper["Q5"] {
		t.Fatalf("got %+v %v", lex, err)
	}
	broken := map[string]string{"d/suffixes.tsv": "x\n", "d/idioms.tsv": "x\n", "d/classes.tsv": "x\n"}
	for name, body := range broken {
		fs := maps.Clone(good)
		fs[name] = &fstest.MapFile{Data: []byte(body)}
		if _, err := FromFS(fs, "d"); err == nil {
			t.Errorf("broken %s: want error", name)
		}
	}
	for _, name := range []string{"d/suffixes.tsv", "d/property_nouns.txt", "d/idioms.tsv", "d/classes.tsv"} {
		fs := maps.Clone(good)
		delete(fs, name)
		if _, err := FromFS(fs, "d"); err == nil {
			t.Errorf("missing %s: want error", name)
		}
	}
}
