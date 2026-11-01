//ff:func feature=text type=parser control=iteration dimension=1
//ff:what TestMakeEojeol — 앞뒤 문장부호를 떼고 안쪽 따옴표를 지우며, 양쪽 따옴표면 인용·쉼표 기록, 문장부호뿐이면 ok 거짓
package text

import "testing"

func TestMakeEojeol(t *testing.T) {
	cases := []struct {
		in   string
		want Eojeol
		ok   bool
	}{
		{"수도이다.", Eojeol{Surface: "수도이다"}, true},
		{"식육목,", Eojeol{Surface: "식육목", Break: true}, true},
		{"\"빛\"을", Eojeol{Surface: "빛을"}, true},
		{"\"phōs\",", Eojeol{Surface: "phōs", Quoted: true, Break: true}, true},
		{"“가수”", Eojeol{Surface: "가수", Quoted: true}, true},
		{"0.61%에", Eojeol{Surface: "0.61%에"}, true},
		{",", Eojeol{Break: true}, false},
	}
	for _, c := range cases {
		got, ok := makeEojeol(c.in)
		if got != c.want || ok != c.ok {
			t.Errorf("%q → %+v %v", c.in, got, ok)
		}
	}
}
