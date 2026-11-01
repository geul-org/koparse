//ff:func feature=stats type=builder control=sequence
//ff:what TestCoverage — 단위가 덮는 어절만 참(범위 밖은 무시)
package stats

import (
	"reflect"
	"testing"

	"github.com/geul-org/koparse/pkg/noun"
)

func TestCoverage(t *testing.T) {
	got := coverage(3, []noun.Unit{{First: 0, Last: 1}, {First: 2, Last: 5}})
	if !reflect.DeepEqual(got, []bool{true, true, true}) || coverage(2, nil)[0] {
		t.Fatal(got)
	}
}
