//ff:type feature=lexicon type=model
//ff:what Suffix — 조사·서술격·용언화 접미 한 항목(형태, 종류, 사슬 자리, 받침 조건, 무게)
package lexicon

// Suffix is one entry of the particle/ending list.
type Suffix struct {
	Form   string  `json:"form"`
	Kind   string  `json:"type"`
	Slot   int     `json:"-"`
	Cond   string  `json:"-"`
	Weight float64 `json:"-"`
}
