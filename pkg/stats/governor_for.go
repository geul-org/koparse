//ff:func feature=stats type=helper control=selection
//ff:what governorFor — 지배어: 의·맨 명사 다음 명사면 N:명사, 주제(은·는)면 문장 끝 서술어, 아니면 다음 서술어
package stats

import "strings"

// governorFor returns the governor of the noun at p and, when it is a
// predicate, its token index (−1 otherwise).
func governorFor(toks []Token, p int, josa string) (string, int) {
	var g int
	switch {
	case strings.HasSuffix(josa, "의"):
		if n := findTag(toks, p+1, 1, "N"); n >= 0 {
			return "N:" + toks[n].Text, -1
		}
		g = findTag(toks, p+1, 1, "V")
	case josa == "∅" && p+1 < len(toks) && toks[p+1].Tag == "N":
		return "N:" + toks[p+1].Text, -1
	case strings.HasSuffix(josa, "은") || strings.HasSuffix(josa, "는"):
		g = findTag(toks, len(toks)-1, -1, "V")
	default:
		g = findTag(toks, p+1, 1, "V")
	}
	if g < 0 {
		return "", -1
	}
	return toks[g].Text, g
}
