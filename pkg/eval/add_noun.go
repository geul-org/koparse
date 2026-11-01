//ff:func feature=eval type=judge control=selection
//ff:what addNoun — 출력 명사 하나를 상태별로 세고, 채택이면 판정 집계로 넘긴다
package eval

import "github.com/geul-org/koparse/pkg/noun"

// addNoun records one output noun.
func addNoun(s *NounScore, ix goldIndex, key string, n noun.Result) {
	s.Nouns++
	switch n.Status {
	case noun.StatusAdopted:
		s.Adopted++
		addAdopted(s, ix, key, n)
	case noun.StatusUndecided:
		s.Undecided++
	default:
		s.Dropped++
		s.DroppedReason[n.Reason]++
	}
}
