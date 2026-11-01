//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what spaceVariants — 어간 그대로와, 세 음절 이상 한글 어간이면 공백을 한 군데 넣은 변형들(R5 공백 정규화)
package noun

import "github.com/geul-org/koparse/pkg/hangul"

// spaceVariants returns stem followed, for Hangul stems of three or more
// syllables, by every one-space variant (발효식품 → 발효 식품).
func spaceVariants(stem string) []string {
	out := []string{stem}
	rs := []rune(stem)
	if !hangul.IsAllHangul(stem) || len(rs) < 3 {
		return out
	}
	for k := 1; k < len(rs); k++ {
		out = append(out, string(rs[:k])+" "+string(rs[k:]))
	}
	return out
}
