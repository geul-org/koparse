//ff:func feature=eval type=helper control=sequence
//ff:what TestSortedKeys — 표본·문장 번호를 수로 정렬(10이 2 뒤)
package eval

import (
	"reflect"
	"testing"

	"github.com/geul-org/koparse/pkg/analyze"
)

func TestSortedKeys(t *testing.T) {
	outs := map[string]analyze.SentenceOut{"10.1": {}, "2.10": {}, "2.2": {}}
	if got := sortedKeys(outs); !reflect.DeepEqual(got, []string{"2.2", "2.10", "10.1"}) {
		t.Fatal(got)
	}
}
