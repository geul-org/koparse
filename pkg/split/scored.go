//ff:type feature=split type=model
//ff:what Scored — 판정을 받은 분리 후보(후보, verdict, 적용도)
package split

// Scored is a candidate with its TAN verdict and applicability.
type Scored struct {
	Candidate     Candidate `json:"candidate"`
	Verdict       float64   `json:"verdict"`
	Applicability float64   `json:"applicability"`
}
