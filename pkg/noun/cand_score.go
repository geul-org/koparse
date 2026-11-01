//ff:type feature=noun type=model
//ff:what candScore — QID 후보 하나의 합친 판정(warrant들의 noisy-OR, 최대 적용도, 살아 있는 규칙)
package noun

// candScore folds the warrants of one candidate.
type candScore struct {
	Miss          float64 // Π(1 − raw) over the candidate's warrants
	Applicability float64
	Rules         []string
}
