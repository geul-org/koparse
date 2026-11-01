//ff:type feature=split type=model
//ff:what Splitter — 어절 분리기(지식 어댑터, 자체 목록, 설정, 어간 후보 수 캐시)
package split

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// Splitter splits eojeols into stem and suffix chain.
type Splitter struct {
	K     knowledge.Knowledge
	Lex   *lexicon.Lexicon
	Cfg   config.Split
	count map[string]int
}
