//ff:type feature=noun type=model
//ff:what Result — 명사 하나의 판정 결과(형태, 채택 QID, 1위 후보, verdict, 적용도, 상태, 지지 규칙, 버린 까닭)
package noun

// Result is the judged noun, as printed in the "nouns" output field.
type Result struct {
	Form          string   `json:"form"`
	QID           string   `json:"qid,omitempty"`
	Best          string   `json:"best,omitempty"`
	Verdict       float64  `json:"verdict"`
	Applicability float64  `json:"applicability"`
	Status        string   `json:"status"`
	Rules         []string `json:"rules,omitempty"`
	Reason        string   `json:"reason,omitempty"`
	Candidates    int      `json:"candidates"`
	Eojeol        [2]int   `json:"eojeol"`
	Piece         bool     `json:"piece,omitempty"`
}
