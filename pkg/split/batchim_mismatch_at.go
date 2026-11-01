//ff:func feature=split type=helper control=iteration dimension=1
//ff:what batchimMismatchAt — 사슬의 어느 접미든 앞말 끝소리와 어긋나면(…으+로는 으로로 봄) 그 접미를 돌려준다
package split

import (
	"strings"

	"github.com/geul-org/koparse/pkg/hangul"
)

// batchimMismatchAt returns the first suffix whose allomorph condition fails
// after the text before it, or "" when all fit.
func batchimMismatchAt(c Candidate) string {
	prev := c.Stem
	for _, s := range c.Suffixes {
		if !s.Accepts(hangul.TailOf(prev)) || (s.Cond == "VL" && strings.HasSuffix(prev, "으")) {
			return s.Form
		}
		prev += s.Form
	}
	return ""
}
