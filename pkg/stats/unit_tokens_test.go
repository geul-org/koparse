//ff:func feature=stats type=builder control=sequence
//ff:what TestUnitTokens — 어절 e에서 시작하는 단위만, 여러 어절 단위는 마지막 어절 번호로
package stats

import (
	"testing"

	"github.com/geul-org/koparse/pkg/noun"
)

func TestUnitTokens(t *testing.T) {
	us := []noun.Unit{{Form: "한강 연안", First: 0, Last: 1}, {Form: "김치", First: 2, Last: 2}}
	got := unitTokens(nil, us, 0, 0)
	if len(got) != 1 || got[0].Text != "한강 연안" || got[0].Eojeol != 1 || got[0].Unit != 0 || got[0].Tag != "N" {
		t.Fatal(got)
	}
}
