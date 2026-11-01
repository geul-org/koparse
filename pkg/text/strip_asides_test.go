//ff:func feature=text type=parser control=iteration dimension=1
//ff:what TestStripAsides — 붙은 괄호, 띄운 괄호, 중첩 괄호, 닫히지 않은 괄호
package text

import (
	"reflect"
	"testing"
)

func TestStripAsides(t *testing.T) {
	cases := []struct {
		in, body string
		asides   []Aside
	}{
		{"서울특별시(-特別市)는 수도", "서울특별시는 수도", []Aside{{Head: "서울특별시", Text: "-特別市"}}},
		{"그리스어 φῶς (빛 [*])와 합성", "그리스어 φῶς와 합성", []Aside{{Head: "φῶς", Text: "빛 [*]"}}},
		{"식육목(Carnivora), 고양이과", "식육목, 고양이과", []Aside{{Head: "식육목", Text: "Carnivora"}}},
		{"열린 (괄호", "열린 (괄호", nil},
	}
	for _, c := range cases {
		body, asides := StripAsides(c.in)
		if body != c.body || !reflect.DeepEqual(asides, c.asides) {
			t.Errorf("%q → %q %v", c.in, body, asides)
		}
	}
}
