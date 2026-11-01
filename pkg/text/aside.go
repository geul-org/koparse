//ff:type feature=text type=model
//ff:what Aside — 괄호 곁가지(앞 낱말과 괄호 안 글). 별칭 후보로 기록한다
package text

// Aside is a parenthetical removed from the sentence body, recorded as an
// alias candidate of the word before it.
type Aside struct {
	Head string `json:"head"`
	Text string `json:"text"`
}
