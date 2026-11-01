//ff:type feature=text type=model
//ff:what Eojeol — 어절 하나(앞뒤 문장부호를 뗀 표면, 따옴표 안 여부, 뒤에 쉼표류가 있는지)
package text

// Eojeol is one space- or punctuation-delimited token.
type Eojeol struct {
	Surface string `json:"surface"`
	// Quoted is true when the token was wrapped in (or touched) quotation marks.
	Quoted bool `json:"quoted,omitempty"`
	// Break is true when a clause-break mark (, ; :) follows the token.
	Break bool `json:"break,omitempty"`
}
