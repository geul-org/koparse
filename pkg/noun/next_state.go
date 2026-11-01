//ff:func feature=noun type=helper control=sequence
//ff:what nextState — 앞 칸에 조각 하나(아는 조각 또는 허용된 접사)를 더한 칸
package noun

// nextState extends prev (ending at j) by one piece. An unknown piece is only
// allowed as an affix.
func nextState(prev segState, j int, known, affixOK bool) segState {
	if !known && !affixOK {
		return segState{}
	}
	s := segState{ok: true, pieces: prev.pieces + 1, affixes: prev.affixes, prev: j, known: known}
	if !known {
		s.affixes++
	}
	return s
}
