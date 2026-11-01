//ff:type feature=eval type=model
//ff:what Sample — 평가 표본 하나(번호, 주제 QID, 원문)
package eval

// Sample is one evaluation document.
type Sample struct {
	ID   int    `json:"id"`
	QID  string `json:"qid"`
	Text string `json:"text"`
}
