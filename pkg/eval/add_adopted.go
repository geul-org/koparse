//ff:func feature=eval type=judge control=iteration dimension=1
//ff:what addAdopted — 채택 명사를 판정해 판정별·규칙별로 세고 목록에 남긴다
package eval

import (
	"strings"

	"github.com/geul-org/koparse/pkg/noun"
)

// addAdopted judges an adopted noun and counts it per judgement and per rule.
func addAdopted(s *NounScore, ix goldIndex, key string, n noun.Result) {
	j := ix.judge(key, n.Form, n.QID)
	s.Judged[j]++
	s.Rows = append(s.Rows, key+"\t"+n.Form+"\t"+n.QID+"\t"+j+"\t"+strings.Join(n.Rules, ","))
	for _, r := range n.Rules {
		c := s.ByRule[r]
		c[judgeIndex(j)]++
		s.ByRule[r] = c
	}
}
