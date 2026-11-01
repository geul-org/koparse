//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what markGeneral — 후보가 여럿이고 고유 분류가 아닌 후보가 하나뿐이면 그 후보에 R4a 표시
package noun

// markGeneral sets General on the only non-proper candidate among several.
func markGeneral(cands []candInfo) {
	only := -1
	for i, c := range cands {
		if c.Proper {
			continue
		}
		if only >= 0 {
			return
		}
		only = i
	}
	if only >= 0 && len(cands) > 1 {
		cands[only].General = true
	}
}
