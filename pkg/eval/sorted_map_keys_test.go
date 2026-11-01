//ff:func feature=eval type=helper control=sequence
//ff:what TestSortedMapKeys — 열쇠 정렬
package eval

import (
	"reflect"
	"testing"
)

func TestSortedMapKeys(t *testing.T) {
	if got := sortedMapKeys(map[string]int{"b": 1, "a": 2}); !reflect.DeepEqual(got, []string{"a", "b"}) {
		t.Fatal(got)
	}
}
