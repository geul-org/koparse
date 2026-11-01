//ff:type feature=text type=model
//ff:what Sentence — 문장 하나(정규화한 원문, 곁가지를 뗀 본문, 어절들, 괄호 곁가지)
package text

// Sentence is one sentence after normalization.
type Sentence struct {
	Text    string   `json:"text"`
	Body    string   `json:"body"`
	Eojeols []Eojeol `json:"eojeols"`
	Asides  []Aside  `json:"asides,omitempty"`
}
