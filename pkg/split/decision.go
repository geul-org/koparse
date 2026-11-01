//ff:type feature=split type=model
//ff:what Decision — 어절 하나의 분리 결정(순위 매긴 후보, 1위, 상태)
package split

// Decision is the judged split of one eojeol. Ranked[0] is the chosen
// candidate; Status is "adopted" or "undecided".
type Decision struct {
	Ranked []Scored `json:"ranked"`
	Status string   `json:"status"`
}
