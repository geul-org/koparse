//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what parseClasses — 분류 목록을 proper·internal 두 집합으로 읽는다
package lexicon

import "fmt"

// parseClasses parses classes.tsv into the proper and internal QID sets.
func parseClasses(s string) (map[string]bool, map[string]bool, error) {
	sets := map[string]map[string]bool{"proper": {}, "internal": {}}
	for _, f := range dataLines(s) {
		set, ok := sets[f[0]]
		if !ok || len(f) < 2 {
			return nil, nil, fmt.Errorf("%v: want proper|internal and a QID", f)
		}
		set[f[1]] = true
	}
	return sets["proper"], sets["internal"], nil
}
