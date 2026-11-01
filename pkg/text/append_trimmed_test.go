//ff:func feature=text type=helper control=sequence
//ff:what TestAppendTrimmed — 공백만 있으면 붙이지 않는다
package text

import (
	"reflect"
	"testing"
)

func TestAppendTrimmed(t *testing.T) {
	got := appendTrimmed(appendTrimmed(nil, "  "), " 가 ")
	if !reflect.DeepEqual(got, []string{"가"}) {
		t.Fatal(got)
	}
}
