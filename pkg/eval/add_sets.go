//ff:func feature=eval type=builder control=iteration dimension=1
//ff:what addSets — 정답 한 줄의 맞는·틀린 QID를 색인 칸에 더한다
package eval

// addSets merges g's good and bad QIDs into m[k].
func addSets(m map[string]nounSets, k string, g GoldNoun) {
	s, ok := m[k]
	if !ok {
		s = nounSets{good: map[string]bool{}, bad: map[string]bool{}}
	}
	for _, q := range g.Good {
		s.good[q] = true
	}
	for _, q := range g.Bad {
		s.bad[q] = true
	}
	m[k] = s
}
