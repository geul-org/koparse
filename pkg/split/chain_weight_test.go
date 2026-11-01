//ff:func feature=split type=helper control=sequence
//ff:what TestChainWeight — 빈 사슬 0, 무게들의 noisy-OR
package split

import (
	"math"
	"testing"

	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestChainWeight(t *testing.T) {
	got := chainWeight([]lexicon.Suffix{{Weight: 0.5}, {Weight: 0.5}})
	if chainWeight(nil) != 0 || math.Abs(got-0.75) > 1e-9 {
		t.Fatal(got)
	}
}
