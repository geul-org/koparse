//ff:func feature=split type=helper control=iteration dimension=1
//ff:what chainWeight — 접미 사슬 무게의 noisy-OR(1 − Π(1 − w))
package split

import "github.com/geul-org/koparse/pkg/lexicon"

// chainWeight combines the weights of a suffix chain by noisy-OR.
func chainWeight(chain []lexicon.Suffix) float64 {
	miss := 1.0
	for _, s := range chain {
		miss *= 1 - s.Weight
	}
	return 1 - miss
}
