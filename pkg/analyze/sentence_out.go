//ff:type feature=analyze type=model
//ff:what SentenceOut — 출력 한 줄(문장, 괄호 곁가지, 어절 분리, 명사 판정, 문맥 통계)
package analyze

import (
	"github.com/geul-org/koparse/pkg/noun"
	"github.com/geul-org/koparse/pkg/stats"
	"github.com/geul-org/koparse/pkg/text"
)

// SentenceOut is one JSON line of output (§5).
type SentenceOut struct {
	Text   string         `json:"text"`
	Asides []text.Aside   `json:"asides,omitempty"`
	Eojeol []EojeolOut    `json:"eojeol"`
	Nouns  []noun.Result  `json:"nouns"`
	Stats  []stats.Record `json:"stats"`
}
