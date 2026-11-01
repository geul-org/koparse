//ff:func feature=stats type=builder control=iteration dimension=1
//ff:what coverage — 명사 단위가 덮는 어절 표시
package stats

import "github.com/geul-org/koparse/pkg/noun"

// coverage marks the eojeols covered by any unit.
func coverage(n int, units []noun.Unit) []bool {
	cov := make([]bool, n)
	for _, u := range units {
		for k := u.First; k <= u.Last && k < n; k++ {
			cov[k] = true
		}
	}
	return cov
}
