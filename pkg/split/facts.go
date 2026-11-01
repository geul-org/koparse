//ff:type feature=split type=model
//ff:what facts — 어절 하나의 판정 재료(후보, 후보별 어간의 지식 후보 수, 따옴표 여부, 목록, 설정)
package split

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// factsKey is the toulmin context key holding *facts.
const factsKey = "split.facts"

// facts is what the split rules read from the toulmin context.
type facts struct {
	cands     []Candidate
	stemCount []int
	quoted    bool
	lex       *lexicon.Lexicon
	cfg       config.Split
}
