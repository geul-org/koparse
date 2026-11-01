//ff:func feature=lexicon type=parser control=iteration dimension=1
//ff:what parseWordList — 한 줄에 낱말 하나인 목록을 집합으로
package lexicon

// parseWordList returns the first field of every data line as a set.
func parseWordList(s string) map[string]bool {
	set := map[string]bool{}
	for _, f := range dataLines(s) {
		set[f[0]] = true
	}
	return set
}
