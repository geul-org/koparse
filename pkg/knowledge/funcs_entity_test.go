//ff:func feature=knowledge type=adapter control=sequence
//ff:what TestFuncsEntity — EntityFn으로 넘긴다
package knowledge

import (
	"context"
	"testing"
)

func TestFuncsEntity(t *testing.T) {
	f := Funcs{EntityFn: func(_ context.Context, q string) (EntityInfo, error) { return EntityInfo{QID: q}, nil }}
	if e, err := f.Entity(context.Background(), "Q2"); err != nil || e.QID != "Q2" {
		t.Fatal(e, err)
	}
}
