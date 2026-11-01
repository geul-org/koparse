//ff:type feature=stats type=model
//ff:what Record — 명사 하나의 문맥 통계(슬롯 묶음: 명사·조사·지배 서술어·부정, 거리 쌍)
package stats

// Record is the context statistics of one noun (adopted, dropped or undecided).
type Record struct {
	Noun     string `json:"noun"`
	QID      string `json:"qid,omitempty"`
	Status   string `json:"status"`
	Josa     string `json:"josa"`
	Governor string `json:"governor"`
	Negated  bool   `json:"negated"`
	Pairs    []Pair `json:"pairs"`
}
