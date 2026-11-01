//ff:func feature=stats type=helper control=iteration dimension=1
//ff:what TestIsPredicate — 어미 글자, 관형형 ㄴ·ㄹ 받침은 서술어, 비한글·다른 받침은 아님
package stats

import "testing"

func TestIsPredicate(t *testing.T) {
	cases := map[string]bool{"달린다": true, "떨어진": true, "배울": true, "가장": false, "km": false, "": false}
	for s, want := range cases {
		if got := isPredicate(s); got != want {
			t.Errorf("isPredicate(%q) = %v", s, got)
		}
	}
}
