//ff:func feature=analyze type=builder control=sequence
//ff:what TestNew — 분리기와 조회 캐시를 갖춘다
package analyze

import (
	"testing"

	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/lexicon"
)

func TestNew(t *testing.T) {
	a := New(nil, &lexicon.Lexicon{}, config.Default())
	if a.splitter == nil || a.memo == nil || a.Cfg.Noun.QR2 != 1 {
		t.Fatal(a)
	}
}
