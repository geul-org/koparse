//ff:func feature=stats type=helper control=iteration dimension=1
//ff:what TestGovernorFor — 의 → 다음 명사(없으면 다음 서술어), 맨 명사 → 바로 뒤 명사, 주제 → 끝 서술어, 그 밖 → 다음 서술어, 서술어 없음
package stats

import "testing"

func TestGovernorFor(t *testing.T) {
	toks := []Token{{Tag: "N", Text: "강"}, {Tag: "N", Text: "하구"}, {Tag: "V", Text: "넓고"}, {Tag: "N", Text: "물"}, {Tag: "V", Text: "맑다"}}
	cases := []struct {
		p    int
		josa string
		want string
		idx  int
	}{
		{0, "의", "N:하구", -1}, {3, "의", "맑다", 4}, {0, "∅", "N:하구", -1},
		{0, "는", "맑다", 4}, {0, "을", "넓고", 2}, {4, "을", "", -1},
	}
	for _, c := range cases {
		g, i := governorFor(toks, c.p, c.josa)
		if g != c.want || i != c.idx {
			t.Errorf("%d %s → %q %d", c.p, c.josa, g, i)
		}
	}
}
