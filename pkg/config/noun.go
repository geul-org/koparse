//ff:type feature=config type=model
//ff:what Noun — 명사 판정 설정
package config

// Noun configures the noun identification judgement (§6.5).
type Noun struct {
	// Qualifiers of the general rules.
	QR1  float64 `json:"q_r1"`
	QR2  float64 `json:"q_r2"`
	QR3  float64 `json:"q_r3"`
	QR4A float64 `json:"q_r4a"`
	QR4B float64 `json:"q_r4b"`
	// Attack forces of the rebuttal and the exceptions.
	ForceClass    float64 `json:"force_class"`
	ForceProperty float64 `json:"force_property"`
	ForceIdiom    float64 `json:"force_idiom"`
	// Adopt is the verdict the first candidate must reach; Margin the gap to the second.
	Adopt  float64 `json:"adopt"`
	Margin float64 `json:"margin"`
}
