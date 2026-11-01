//ff:func feature=text type=parser control=sequence
//ff:what TestParse — 문장 둘, 곁가지 기록, 어절 분할
package text

import "testing"

func TestParse(t *testing.T) {
	got := Parse("김치(沈菜)는 음식이다. 맛있다")
	if len(got) != 2 || got[0].Body != "김치는 음식이다." || len(got[0].Asides) != 1 || len(got[0].Eojeols) != 2 || got[1].Eojeols[0].Surface != "맛있다" {
		t.Fatalf("%+v", got)
	}
}
