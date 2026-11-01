//ff:type feature=split type=model
//ff:what Candidate — 어절 분리 후보(어간 + 조사·어미 사슬). 사슬이 비면 통째 후보
package split

import "github.com/geul-org/koparse/pkg/lexicon"

// Candidate is one way to split an eojeol: a stem followed by a suffix chain.
// An empty chain is the whole-eojeol candidate.
type Candidate struct {
	Stem     string           `json:"stem"`
	Suffixes []lexicon.Suffix `json:"suffixes,omitempty"`
}
