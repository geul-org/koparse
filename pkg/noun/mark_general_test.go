//ff:func feature=noun type=helper control=iteration dimension=2
//ff:what TestMarkGeneral — 일반 후보가 하나뿐일 때만(후보 하나·일반 둘·일반 없음은 표시 없음)
package noun

import "testing"

func TestMarkGeneral(t *testing.T) {
	one := []candInfo{{Proper: true}, {}}
	markGeneral(one)
	if !one[1].General || one[0].General {
		t.Fatal(one)
	}
	for _, cs := range [][]candInfo{{{}}, {{}, {}}, {{Proper: true}, {Proper: true}}} {
		markGeneral(cs)
		for _, c := range cs {
			if c.General {
				t.Fatal(cs)
			}
		}
	}
}
