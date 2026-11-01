//ff:func feature=lexicon type=parser control=sequence
//ff:what TestDataLines — 주석·빈 줄을 빼고 줄 끝 주석을 지우며 필드 공백을 다듬는다
package lexicon

import (
	"reflect"
	"testing"
)

func TestDataLines(t *testing.T) {
	got := dataLines("# head\n\n a \t b # tail\nc\n")
	want := [][]string{{"a", "b"}, {"c"}}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
