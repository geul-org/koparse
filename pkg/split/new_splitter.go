//ff:func feature=split type=builder control=sequence
//ff:what NewSplitter — 분리기를 만든다
package split

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// NewSplitter returns a Splitter with an empty stem-count cache.
func NewSplitter(k knowledge.Knowledge, lex *lexicon.Lexicon, cfg config.Split) *Splitter {
	return &Splitter{K: k, Lex: lex, Cfg: cfg, count: map[string]int{}}
}
