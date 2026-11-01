//ff:type feature=noun type=model
//ff:what Memo — 지식 조회 캐시(표면→위키 내부 문서를 뺀 후보, QID→개체 요약)
package noun

import (
	"github.com/geul-org/koparse/pkg/knowledge"
	"github.com/geul-org/koparse/pkg/lexicon"
)

// Memo caches knowledge lookups for one analysis run (§4.4: the same word
// repeats within a document).
type Memo struct {
	K     knowledge.Knowledge
	Lex   *lexicon.Lexicon
	cands map[string]lookupResult
	ents  map[string]knowledge.EntityInfo
}
