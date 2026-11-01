//ff:func feature=stats type=helper control=iteration dimension=1
//ff:what findTag — from부터 앞으로 품사가 tag인 첫 토큰 번호(step −1이면 끝에서부터 거꾸로), 없으면 −1
package stats

// findTag scans toks from index from by step (+1 forward, −1 backward) and
// returns the first index with the tag, or −1.
func findTag(toks []Token, from, step int, tag string) int {
	for i := from; i >= 0 && i < len(toks); i += step {
		if toks[i].Tag == tag {
			return i
		}
	}
	return -1
}
