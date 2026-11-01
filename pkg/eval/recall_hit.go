//ff:func feature=eval type=judge control=iteration dimension=1
//ff:what recallHit — 정답 명사 한 번에 맞는 채택 명사가 문장 안에 있으면 그 번호를 쓰고 참
package eval

import "github.com/geul-org/koparse/pkg/noun"

// recallHit marks and reports the first unused adopted noun of the sentence
// with the gold form and a good QID.
func recallHit(nouns []noun.Result, used []bool, g GoldNoun) bool {
	good := map[string]bool{}
	for _, q := range g.Good {
		good[q] = true
	}
	for i, n := range nouns {
		if !used[i] && n.Status == noun.StatusAdopted && n.Form == g.Noun && good[n.QID] {
			used[i] = true
			return true
		}
	}
	return false
}
