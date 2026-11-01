//ff:func feature=analyze type=helper control=sequence
//ff:what TestCoveredEojeols — 단위 범위의 어절 번호
package analyze

import (
	"testing"

	"github.com/geul-org/koparse/pkg/noun"
)

func TestCoveredEojeols(t *testing.T) {
	got := coveredEojeols([]noun.Unit{{First: 1, Last: 2}})
	if len(got) != 2 || !got[1] || !got[2] {
		t.Fatal(got)
	}
}
