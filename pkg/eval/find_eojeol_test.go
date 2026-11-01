//ff:func feature=eval type=helper control=sequence
//ff:what TestFindEojeol — 안 쓴 어절 중 같은 표면, 없으면 −1
package eval

import (
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
)

func TestFindEojeol(t *testing.T) {
	es := []analyze.EojeolOut{{Surface: "가"}, {Surface: "가"}}
	if findEojeol(es, []bool{true, false}, "가") != 1 || findEojeol(es, []bool{true, true}, "가") != -1 {
		t.Fatal("findEojeol")
	}
}
