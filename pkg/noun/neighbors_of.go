//ff:func feature=noun type=helper control=iteration dimension=1
//ff:what neighborsOf — 채택된 명사들의 QID·분류로 이웃 집합을 만든다
package noun

// neighborsOf builds the neighbor set from adopted results.
func neighborsOf(units []Unit, results []Result) map[string]bool {
	nb := map[string]bool{}
	for i, r := range results {
		if r.Status == StatusAdopted {
			addNeighbor(nb, units[i], r.QID)
		}
	}
	return nb
}
