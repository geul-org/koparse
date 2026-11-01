//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what markNeighbor — 일반 후보 중 이웃 집합과 QID·P31·P279가 겹치는 것이 하나뿐이면 그 후보에 R4b 표시
package noun

// markNeighbor sets Neighbor on the only non-proper candidate that shares a
// QID or class with nb.
func markNeighbor(cands []candInfo, u Unit, nb map[string]bool) {
	only := -1
	for i, c := range cands {
		e := u.Infos[i]
		if c.Proper || !(nb[c.QID] || anyIn(e.P31, nb) || anyIn(e.P279, nb)) {
			continue
		}
		if only >= 0 {
			return
		}
		only = i
	}
	if only >= 0 {
		cands[only].Neighbor = true
	}
}
