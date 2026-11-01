//ff:type feature=analyze type=model
//ff:what Part — 분리 결과 조각(형태, 종류 noun·word·particle·copula·verbal)
package analyze

// Part is one piece of an eojeol split.
type Part struct {
	Form string `json:"form"`
	Type string `json:"type"`
}
