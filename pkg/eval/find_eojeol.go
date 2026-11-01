//ff:func feature=eval type=helper control=iteration dimension=1
//ff:what findEojeol — 아직 안 쓴 출력 어절 중 표면이 같은 첫 번호(없으면 −1)
package eval

import "github.com/geul-org/koparse/pkg/analyze"

// findEojeol returns the first unused output eojeol with the given surface.
func findEojeol(es []analyze.EojeolOut, used []bool, surface string) int {
	for i, e := range es {
		if !used[i] && e.Surface == surface {
			return i
		}
	}
	return -1
}
