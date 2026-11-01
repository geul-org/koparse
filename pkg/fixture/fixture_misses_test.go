//ff:func feature=fixture type=adapter control=sequence
//ff:what TestFixtureMisses — 놓친 표면을 정렬해 돌려준다
package fixture

import (
	"context"
	"reflect"
	"testing"
)

func TestFixtureMisses(t *testing.T) {
	f := New(nil, nil, nil)
	f.Candidates(context.Background(), "나", "ko")
	f.Candidates(context.Background(), "가", "ko")
	f.Candidates(context.Background(), "가", "ko")
	if got := f.Misses(); !reflect.DeepEqual(got, []string{"ko:가", "ko:나"}) {
		t.Fatal(got)
	}
}
