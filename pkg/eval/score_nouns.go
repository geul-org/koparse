//ff:func feature=eval type=judge control=iteration dimension=2
//ff:what ScoreNouns — 출력 명사 모두를 정답과 손 판정으로 집계하고 재현을 센다
package eval

import "github.com/geul-org/koparse/pkg/analyze"

// ScoreNouns aggregates the noun metrics over every output sentence.
func ScoreNouns(outs map[string]analyze.SentenceOut, gold []GoldNoun, extra map[string]string) NounScore {
	s := NounScore{DroppedReason: map[string]int{}, Judged: map[string]int{}, ByRule: map[string][3]int{}}
	ix := newGoldIndex(gold, extra)
	for _, key := range sortedKeys(outs) {
		for _, n := range outs[key].Nouns {
			addNoun(&s, ix, key, n)
		}
	}
	countRecall(&s, outs, gold)
	return s
}
