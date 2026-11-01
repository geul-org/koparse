//ff:type feature=analyze type=model
//ff:what Analyzer — 전체 처리기(지식 어댑터, 자체 목록, 설정, 분리기, 조회 캐시, 후보 순위 출력 여부)
package analyze

import (
	"github.com/geul-org/koparse/pkg/config"
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
	"github.com/geul-org/koparse/pkg/noun"
	"github.com/geul-org/koparse/pkg/split"
)

// Analyzer runs the whole pipeline (§6) over one document at a time.
type Analyzer struct {
	K        knowledge.Knowledge
	Lex      *lexicon.Lexicon
	Cfg      config.Config
	Trace    bool
	splitter *split.Splitter
	memo     *noun.Memo
}
