//ff:func feature=split type=helper control=sequence
//ff:what TestFactsOf — 문맥에 넣은 재료를 그대로 꺼낸다
package split

import (
	"testing"

	"github.com/park-jun-woo/toulmin/go/pkg/toulmin"
)

func TestFactsOf(t *testing.T) {
	f := &facts{quoted: true}
	ctx := toulmin.NewContext()
	ctx.Set(factsKey, f)
	if factsOf(ctx) != f {
		t.Fatal("facts")
	}
}
