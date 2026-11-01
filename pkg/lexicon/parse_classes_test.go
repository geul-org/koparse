//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what TestParseClasses — proper·internal 두 집합, 모르는 묶음·QID 없는 줄은 오류
package lexicon

import "testing"

func TestParseClasses(t *testing.T) {
	p, i, err := parseClasses("proper\tQ5\t# human\ninternal\tQ4167410\n")
	if err != nil || !p["Q5"] || !i["Q4167410"] || len(p) != 1 || len(i) != 1 {
		t.Fatalf("got %v %v %v", p, i, err)
	}
	for _, bad := range []string{"other\tQ1\n", "proper\n"} {
		if _, _, err := parseClasses(bad); err == nil {
			t.Errorf("%q: want error", bad)
		}
	}
}
