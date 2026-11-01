//ff:func feature=lexicon type=rule control=iteration dimension=1
//ff:what ChainValid — 접미 사슬의 순서가 자리 규칙(오름차순, 자리 4 이상은 끝)에 맞는지
package lexicon

// ChainValid reports whether the suffixes may follow each other in this order:
// slots strictly increase and nothing follows a slot ≥ 4, unless the pair is
// listed with @allow.
func (l *Lexicon) ChainValid(chain []Suffix) bool {
	for i := 1; i < len(chain); i++ {
		a, b := chain[i-1], chain[i]
		if l.Allow[a.Form+"\t"+b.Form] {
			continue
		}
		if a.Slot >= 4 || b.Slot <= a.Slot {
			return false
		}
	}
	return true
}
