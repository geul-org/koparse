//ff:func feature=split type=helper control=iteration dimension=1
//ff:what hasHangul — 한글 음절이 하나라도 있는지
package split

import "github.com/geul-org/koparse/pkg/hangul"

// hasHangul reports whether s contains a Hangul syllable.
func hasHangul(s string) bool {
	for _, r := range s {
		if hangul.IsSyllable(r) {
			return true
		}
	}
	return false
}
