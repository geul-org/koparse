//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what particleFits — 관용 틀의 조사 조건(* 아무것, - 없음, 형태)이 붙은 사슬에 맞는지
package noun

import "slices"

// particleFits reports whether chain satisfies any of the idiom particles.
func particleFits(ps []string, chain []string) bool {
	for _, p := range ps {
		if p == "*" || (p == "-" && len(chain) == 0) || slices.Contains(chain, p) {
			return true
		}
	}
	return false
}
