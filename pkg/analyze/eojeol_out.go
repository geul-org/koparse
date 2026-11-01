//ff:type feature=analyze type=model
//ff:what EojeolOut — 출력의 어절 하나(표면, 분리 조각, verdict, 상태, 요청 시 후보 순위)
package analyze

import "github.com/geul-org/koparse/pkg/split"

// EojeolOut is one eojeol of the JSON-line output.
type EojeolOut struct {
	Surface string         `json:"surface"`
	Split   []Part         `json:"split"`
	Verdict float64        `json:"verdict"`
	Status  string         `json:"status"`
	Ranked  []split.Scored `json:"ranked,omitempty"`
}
