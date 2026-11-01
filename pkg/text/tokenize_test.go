//ff:func feature=text type=parser control=iteration dimension=1
//ff:what TestTokenize — 공백·가운뎃점으로 나누고 쉼표 표시, 수의 쉼표는 둔다
package text

import "testing"

func TestTokenize(t *testing.T) {
	got := Tokenize("5,066명으로 경기도·인천광역시와 함께, 수도권을")
	want := []Eojeol{{Surface: "5,066명으로"}, {Surface: "경기도"}, {Surface: "인천광역시와"}, {Surface: "함께", Break: true}, {Surface: "수도권을"}}
	if len(got) != len(want) {
		t.Fatalf("%+v", got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("%d: %+v, want %+v", i, got[i], want[i])
		}
	}
}
