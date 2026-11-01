//ff:func feature=noun type=builder control=sequence
//ff:what NewMemo — 빈 조회 캐시를 만든다
package noun

import (
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// NewMemo returns an empty lookup cache over k.
func NewMemo(k knowledge.Knowledge, lex *lexicon.Lexicon) *Memo {
	return &Memo{K: k, Lex: lex, cands: map[string]lookupResult{}, ents: map[string]knowledge.EntityInfo{}}
}
