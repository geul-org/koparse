//ff:func feature=split type=builder control=sequence
//ff:what TestNewSplitter — 필드와 빈 캐시
package split

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/fixture"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestNewSplitter(t *testing.T) {
	fx := fixture.New(nil, nil, nil)
	lex := &lexicon.Lexicon{}
	s := NewSplitter(fx, lex, config.Split{Margin: 0.1})
	if s.K != fx || s.Lex != lex || s.Cfg.Margin != 0.1 || s.count == nil {
		t.Fatal(s)
	}
}
