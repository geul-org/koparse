//ff:type feature=noun type=model
//ff:what Unit — 명사 단위(표면, 일치한 형태, 어절 범위, 붙은 조사, 지배 서술어, 복합어 조각 여부, 후보와 개체 요약)
package noun

import "github.com/geul-org/koparse/pkg/knowledge"

// Unit is one noun found in a sentence: a stem, a multi-eojeol span, or a
// piece of a compound stem.
type Unit struct {
	Form     string
	Matched  string // the looked-up form that matched (differs from Form for R5 spacing)
	First    int    // first eojeol index
	Last     int    // last eojeol index
	Piece    bool   // a piece of a compound stem
	Final    bool   // ends where the eojeol's stem ends (carries the eojeol's suffixes)
	Chain    []string
	Governor string // the eojeol's own copula/verbal suffix or the next eojeol (idiom check)
	Cands    []knowledge.Candidate
	Infos    []knowledge.EntityInfo
}
