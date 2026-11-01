//ff:func feature=lexicon type=parser control=sequence
//ff:what TestParseIdioms — 세 필드를 | 로 나누고, 필드가 모자라면 오류
package lexicon

import (
	"reflect"
	"testing"
)

func TestParseIdioms(t *testing.T) {
	got, err := parseIdioms("*\t으로|로\t삼|하\n")
	want := []Idiom{{Noun: "*", Particles: []string{"으로", "로"}, Governors: []string{"삼", "하"}}}
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v %v", got, err)
	}
	if _, err := parseIdioms("자리\t-\n"); err == nil {
		t.Fatal("want error")
	}
}
