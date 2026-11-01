//ff:type feature=noun type=model
//ff:what nounSpec — 규칙이 맡은 QID 후보 번호(toulmin Spec)
package noun

// nounSpec attaches a QID candidate index to a rule node.
type nounSpec struct {
	Cand int `json:"cand"`
}
