//ff:func feature=hangul type=helper control=iteration dimension=1
//ff:what TestIsAllHangul — 빈 문자열·섞인 문자열은 거짓
package hangul

import "testing"

func TestIsAllHangul(t *testing.T) {
	cases := map[string]bool{"": false, "김치": true, "km2로": false, "김치A": false}
	for s, want := range cases {
		if got := IsAllHangul(s); got != want {
			t.Errorf("IsAllHangul(%q) = %v, want %v", s, got, want)
		}
	}
}
