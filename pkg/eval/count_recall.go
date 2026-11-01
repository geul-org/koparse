//ff:func feature=eval type=judge control=iteration dimension=1
//ff:what countRecall — 맞는 QID가 있는 정답 명사 중 같은 문장에서 그 QID로 채택된 수
package eval

import "github.com/geul-org/koparse/pkg/analyze"

// countRecall counts the gold nouns (with a good QID) recovered by an adopted
// noun of the same sentence, form and QID.
func countRecall(s *NounScore, outs map[string]analyze.SentenceOut, gold []GoldNoun) {
	used := map[string][]bool{}
	for _, g := range gold {
		if len(g.Good) == 0 {
			continue
		}
		s.GoldGood++
		if used[g.Key] == nil {
			used[g.Key] = make([]bool, len(outs[g.Key].Nouns))
		}
		if recallHit(outs[g.Key].Nouns, used[g.Key], g) {
			s.Recalled++
		}
	}
}
