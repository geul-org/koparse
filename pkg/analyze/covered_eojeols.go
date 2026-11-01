//ff:func feature=analyze type=helper control=iteration dimension=1
//ff:what coveredEojeols — 명사 단위가 덮는 어절 번호 집합
package analyze

import "github.com/geul-org/koparse/pkg/noun"

// coveredEojeols returns the set of eojeol indexes covered by units.
func coveredEojeols(units []noun.Unit) map[int]bool {
	set := map[int]bool{}
	for _, u := range units {
		for k := u.First; k <= u.Last; k++ {
			set[k] = true
		}
	}
	return set
}
