//ff:type feature=split type=model
//ff:what candSpec — 규칙이 맡은 분리 후보 번호(toulmin Spec)
package split

// candSpec attaches a candidate index to a rule node.
type candSpec struct {
	Idx int `json:"idx"`
}
