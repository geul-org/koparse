//ff:func feature=text type=parser control=iteration dimension=1
//ff:what SplitSentences — 마침표·물음표·느낌표·줄바꿈으로 문장을 나눈다(괄호 안은 나누지 않음)
package text

// SplitSentences splits s at sentence ends outside brackets.
func SplitSentences(s string) []string {
	rs := []rune(s)
	var out []string
	start, depth := 0, 0
	for i, r := range rs {
		depth = parenDepth(depth, r)
		if depth == 0 && isSentenceEnd(rs, i) {
			out = appendTrimmed(out, string(rs[start:i+1]))
			start = i + 1
		}
	}
	return appendTrimmed(out, string(rs[start:]))
}
