//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what addNeighbor — 채택된 QID와 그 P31·P279를 이웃 집합에 넣는다
package noun

// addNeighbor adds qid and its P31·P279 classes (from u's summaries) to nb.
func addNeighbor(nb map[string]bool, u Unit, qid string) {
	nb[qid] = true
	for i, c := range u.Cands {
		if c.QID != qid {
			continue
		}
		for _, v := range append(append([]string{}, u.Infos[i].P31...), u.Infos[i].P279...) {
			nb[v] = true
		}
	}
}
