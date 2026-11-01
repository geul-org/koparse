//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what parseIdioms — 관용 틀(명사, 조사|…, 서술어 앞부분|…)을 읽는다
package lexicon

import (
	"fmt"
	"strings"
)

// parseIdioms parses idioms.tsv.
func parseIdioms(s string) ([]Idiom, error) {
	var out []Idiom
	for _, f := range dataLines(s) {
		if len(f) != 3 {
			return nil, fmt.Errorf("%v: want 3 fields", f)
		}
		out = append(out, Idiom{Noun: f[0], Particles: strings.Split(f[1], "|"), Governors: strings.Split(f[2], "|")})
	}
	return out, nil
}
