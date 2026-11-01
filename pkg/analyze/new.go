//ff:func feature=analyze type=builder control=sequence
//ff:what New — 처리기를 만든다
package analyze

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
	"github.com/geul-org/koparse/pkg/split"
)

// New returns an Analyzer over k.
func New(k knowledge.Knowledge, lex *lexicon.Lexicon, cfg config.Config) *Analyzer {
	return &Analyzer{K: k, Lex: lex, Cfg: cfg, splitter: split.NewSplitter(k, lex, cfg.Split), memo: noun.NewMemo(k, lex)}
}
