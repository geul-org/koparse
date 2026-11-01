//ff:func feature=stats type=helper control=sequence
//ff:what TestNegatedAt — 어절 안의 않·못, 다음 어절 머리의 않, 부정 없음
package stats

import (
	"testing"

	"github.com/geul-org/koparse/pkg/noun"
)

func TestNegatedAt(t *testing.T) {
	w := []noun.Word{{Surface: "먹지"}, {Surface: "않는다"}, {Surface: "못한다"}, {Surface: "먹는다"}}
	if !negatedAt(w, 0) || !negatedAt(w, 1) || !negatedAt(w, 2) || negatedAt(w, 3) {
		t.Fatal("negatedAt")
	}
}
