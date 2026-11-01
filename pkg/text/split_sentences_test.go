//ff:func feature=text type=parser control=sequence
//ff:what TestSplitSentences — 마침표·물음표·줄바꿈으로 나누고 소수점·괄호 안 마침표는 두며 끝 문장도 넣는다
package text

import (
	"reflect"
	"testing"
)

func TestSplitSentences(t *testing.T) {
	got := SplitSentences("면적은 605.25 km2다. 왜? (가. 나) 끝\n마지막")
	want := []string{"면적은 605.25 km2다.", "왜?", "(가. 나) 끝", "마지막"}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %q", got)
	}
}
